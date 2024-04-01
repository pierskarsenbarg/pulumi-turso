// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("turso");

/**
 * Your Turso API token
 */
export declare const apiToken: string;
Object.defineProperty(exports, "apiToken", {
    get() {
        return __config.get("apiToken") ?? (utilities.getEnv("TURSO_APITOKEN") || "");
    },
    enumerable: true,
});

/**
 * Organisation name
 */
export declare const organizationName: string;
Object.defineProperty(exports, "organizationName", {
    get() {
        return __config.get("organizationName") ?? (utilities.getEnv("TURSO_ORGANISATIONNAME") || "");
    },
    enumerable: true,
});

