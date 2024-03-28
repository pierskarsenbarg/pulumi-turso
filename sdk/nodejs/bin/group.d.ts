import * as pulumi from "@pulumi/pulumi";
export declare class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Group;
    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Group;
    /**
     * The current libSQL server version the databases in that group are running.
     */
    readonly dbVersion: pulumi.Output<string>;
    /**
     * The group universal unique identifier (UUID).
     */
    readonly groupId: pulumi.Output<string>;
    /**
     * An array of location keys the group is located.
     */
    readonly locations: pulumi.Output<string[] | undefined>;
    /**
     * The group name, unique across your organization.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The name of the organization or user.
     */
    readonly organization: pulumi.Output<string>;
    /**
     * The primary location key.
     */
    readonly primaryLocation: pulumi.Output<string>;
    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupArgs, opts?: pulumi.CustomResourceOptions);
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
