"use strict";
// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
exports.Provider = void 0;
const pulumi = require("@pulumi/pulumi");
const utilities = require("./utilities");
class Provider extends pulumi.ProviderResource {
    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj) {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }
    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name, args, opts) {
        var _a, _b;
        let resourceInputs = {};
        opts = opts || {};
        {
            resourceInputs["apiToken"] = (_a = ((args === null || args === void 0 ? void 0 : args.apiToken) ? pulumi.secret(args.apiToken) : undefined)) !== null && _a !== void 0 ? _a : (utilities.getEnv("TURSO_APITOKEN") || "");
            resourceInputs["organizationName"] = (_b = (args ? args.organizationName : undefined)) !== null && _b !== void 0 ? _b : (utilities.getEnv("TURSO_ORGANISATIONNAME") || "");
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["apiToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}
exports.Provider = Provider;
/** @internal */
Provider.__pulumiType = 'turso';
//# sourceMappingURL=provider.js.map