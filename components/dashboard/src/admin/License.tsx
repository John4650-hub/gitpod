/**
 * Copyright (c) 2022 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License-AGPL.txt in the project root for license information.
 */

import { PageWithSubMenu } from "../components/PageWithSubMenu";
import { adminMenu } from "./admin-menu";

import { LicenseContext } from "../license-context";
import { ReactElement, useContext, useEffect } from "react";
import { getGitpodService } from "../service/service";

import { ReactComponent as Alert } from "../images/exclamation.svg";
import { ReactComponent as Success } from "../images/check-circle.svg";
import { LicenseInfo } from "@gitpod/gitpod-protocol";
import { ReactComponent as XSvg } from "../images/x.svg";
import { ReactComponent as CheckSvg } from "../images/check.svg";
import SolidCard from "../components/SolidCard";
import Card from "../components/Card";

export default function License() {
    const { license, setLicense } = useContext(LicenseContext);

    useEffect(() => {
        if (isGitpodIo()) {
            return; // temporarily disable to avoid hight CPU on the DB
        }
        (async () => {
            const data = await getGitpodService().server.adminGetLicense();
            setLicense(data);
        })();
    }, []);

    const featureList = license?.enabledFeatures;
    const features = license?.features;

    // if user seats is 0, it means that there is user limit in the license
    const userLimit = license?.seats == 0 ? "Unlimited" : license?.seats;

    const [licenseLevel, paid, statusMessage] = license ? getSubscriptionLevel(license) : defaultMessage();

    return (
        <div>
            <PageWithSubMenu
                subMenu={adminMenu}
                title="License"
                subtitle="License associated with your Gitpod Installation"
            >
                <div className="flex flex-row space-x-4">
                    <Card className="bg-gray-800 dark:bg-gray-100 text-gray-300 dark:text-gray-500">
                        {licenseLevel}
                        {paid}
                        <div className="pt-4 font-semibold">Available features:</div>
                        <div className="flex flex-col pt-1">
                            {features &&
                                features.map((feat: string) => (
                                    <span className="inline-flex">
                                        {featureList?.includes(feat) ? (
                                            <CheckSvg fill="currentColor" className="w-4 h-4 mt-1" />
                                        ) : (
                                            <XSvg fill="currentColor" className="pt-1 h-3 mt-1" />
                                        )}
                                        <span>{capitalizeInitials(feat)}</span>
                                    </span>
                                ))}
                        </div>
                    </Card>
                    <SolidCard>
                        <div className="text-gray-600 dark:text-gray-200 py-4 flex-row flex font-semibold items-center">
                            {statusMessage}
                        </div>
                        <p className="dark:text-gray-500 font-semibold">Registered Users</p>
                        <span className="dark:text-gray-300 pt-1 text-lg">{license?.userCount || 0}</span>
                        <span className="dark:text-gray-500 text-gray-400 pt-1 text-lg"> / {userLimit} </span>
                        <p className="dark:text-gray-500 pt-2 font-semibold">License Type</p>
                        <h4 className="dark:text-gray-300 pt-1 text-lg">{capitalizeInitials(license?.type || "")}</h4>
                        <div className="flex justify-end bottom-2 relative">
                            <button
                                type="button"
                                onClick={(e) => {
                                    e.preventDefault();
                                    window.location.href = "https://www.gitpod.io/self-hosted";
                                }}
                            >
                                <div className="font-semibold">Compare Plans</div>
                            </button>
                        </div>
                    </SolidCard>
                </div>
            </PageWithSubMenu>
        </div>
    );
}

function capitalizeInitials(str: string): string {
    return str
        .split("-")
        .map((item) => {
            return item.charAt(0).toUpperCase() + item.slice(1);
        })
        .join(" ");
}

function getSubscriptionLevel(license: LicenseInfo): ReactElement[] {
    switch (license.plan) {
        case "prod":
        case "trial":
            return professionalPlan(license.userCount || 0, license.seats, license.plan == "trial", license.validUntil);
        case "community":
            return communityPlan(license.userCount || 0, license.seats, license.fallbackAllowed);
        default: {
            return defaultMessage();
        }
    }
}

function licenseLevel(level: string): ReactElement {
    return <p className="text-white dark:text-black font-bold pt-4"> {level}</p>;
}

function additionalLicenseInfo(data: string): ReactElement {
    return <p className="dark:text-gray-500">{data}</p>;
}

function defaultMessage(): ReactElement[] {
    const alertMessage = () => {
        return (
            <>
                <div>Inactive or unknown license</div>
                <div className="flex justify-right my-4 mx-1">
                    <Alert fill="grey" className="h-8 w-8" />
                </div>
            </>
        );
    };

    return [licenseLevel("Inactive"), additionalLicenseInfo("Free"), alertMessage()];
}

function professionalPlan(userCount: number, seats: number, trial: boolean, validUntil: string): ReactElement[] {
    const alertMessage = (aboveLimit: boolean) => {
        return aboveLimit ? (
            <>
                <div className="text-red-600">You have exceeded the usage limit.</div>
                <div className="flex justify-right my-4 mx-1">
                    <Alert fill="red" className="h-6 w-6" />
                </div>
            </>
        ) : (
            <>
                <div className="text-green-600">You have an active professional license.</div>
                <div className="flex justify-right my-4 mx-1">
                    <Success fill="green" className="h-8 w-8" />
                </div>
            </>
        );
    };

    const aboveLimit: boolean = userCount > seats;

    const licenseTitle = () => {
        if (trial) {
            const expDate = new Date(validUntil);
            if (typeof expDate.getTime !== "function") {
                return additionalLicenseInfo("Trial");
            } else {
                return additionalLicenseInfo("Trial expires on " + expDate.toLocaleDateString());
            }
        } else {
            return additionalLicenseInfo("Paid");
        }
    };

    return [licenseLevel("Professional"), licenseTitle(), alertMessage(aboveLimit)];
}

function communityPlan(userCount: number, seats: number, fallbackAllowed: boolean): ReactElement[] {
    const alertMessage = (aboveLimit: boolean) => {
        if (aboveLimit) {
            return fallbackAllowed ? (
                <>
                    <div>No active license. You are using community edition.</div>
                    <div className="flex justify-right my-4 mx-1">
                        <Success fill="grey" className="h-8 w-8" />
                    </div>
                </>
            ) : (
                <>
                    <div className="text-red-600">No active license. You have exceeded the usage limit.</div>
                    <div className="flex justify-right my-4 mx-1">
                        <Alert fill="red" className="h-8 w-8" />
                    </div>
                </>
            );
        } else {
            return (
                <>
                    <div>You are using the free community edition.</div>
                    <div className="flex justify-right my-4 mx-1">
                        <Success fill="green" className="h-8 w-8" />
                    </div>
                </>
            );
        }
    };

    const aboveLimit: boolean = userCount > seats;

    return [licenseLevel("Community"), additionalLicenseInfo("Free"), alertMessage(aboveLimit)];
}

function isGitpodIo() {
    return window.location.hostname === "gitpod.io" || window.location.hostname === "gitpod-staging.com";
}
