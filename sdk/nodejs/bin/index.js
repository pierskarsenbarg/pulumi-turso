"use strict";
// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
exports.config = exports.Provider = exports.Group = exports.Database = void 0;
const pulumi = require("@pulumi/pulumi");
const utilities = require("./utilities");
exports.Database = null;
utilities.lazyLoad(exports, ["Database"], () => require("./database"));
exports.Group = null;
utilities.lazyLoad(exports, ["Group"], () => require("./group"));
exports.Provider = null;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));
// Export sub-modules:
const config = require("./config");
exports.config = config;
const _module = {
    version: utilities.getVersion(),
    construct: (name, type, urn) => {
        switch (type) {
            case "turso:index:Database":
                return new exports.Database(name, undefined, { urn });
            case "turso:index:Group":
                return new exports.Group(name, undefined, { urn });
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("turso", "index", _module);
pulumi.runtime.registerResourcePackage("turso", {
    version: utilities.getVersion(),
    constructProvider: (name, type, urn) => {
        if (type !== "pulumi:providers:turso") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new exports.Provider(name, undefined, { urn });
    },
});
//# sourceMappingURL=index.js.map