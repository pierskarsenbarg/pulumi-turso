// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Group {
        return new Group(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'turso:index:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    /**
     * The current libSQL server version the databases in that group are running.
     */
    public /*out*/ readonly dbVersion!: pulumi.Output<string>;
    /**
     * The group universal unique identifier (UUID).
     */
    public /*out*/ readonly groupId!: pulumi.Output<string>;
    /**
     * An array of location keys the group is located.
     */
    public readonly locations!: pulumi.Output<string[] | undefined>;
    /**
     * The group name, unique across your organization.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the organization or user.
     */
    public readonly organization!: pulumi.Output<string>;
    /**
     * The primary location key.
     */
    public readonly primaryLocation!: pulumi.Output<string>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.organization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organization'");
            }
            if ((!args || args.primaryLocation === undefined) && !opts.urn) {
                throw new Error("Missing required property 'primaryLocation'");
            }
            resourceInputs["locations"] = args ? args.locations : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["primaryLocation"] = args ? args.primaryLocation : undefined;
            resourceInputs["dbVersion"] = undefined /*out*/;
            resourceInputs["groupId"] = undefined /*out*/;
        } else {
            resourceInputs["dbVersion"] = undefined /*out*/;
            resourceInputs["groupId"] = undefined /*out*/;
            resourceInputs["locations"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["organization"] = undefined /*out*/;
            resourceInputs["primaryLocation"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Group.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    /**
     * An array of location keys the group is located.
     */
    locations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the new group.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the organization or user.
     */
    organization: pulumi.Input<string>;
    /**
     * The primary location key for the new group.
     */
    primaryLocation: pulumi.Input<string>;
}