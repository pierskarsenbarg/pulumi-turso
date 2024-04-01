import { $ } from "bun";

await $`git update-index -q --refresh`;

const { stdout } = await $`git diff-files`.quiet();

if (stdout.length !== 0) {
    console.log(`Working directory dirty.`);
    await $`git status`
    await $`git diff`
    await $`exit 1`
}