import { $ } from "bun";

await $`git update-index -q --refresh`;

const { stdout, stderr, exitCode } = await $`git diff-files`.quiet().nothrow();

console.log(`exit code: ${exitCode}`);

if (exitCode !== 0) {
    console.log(`Working directory dirty.`);
    await $`git status`
    await $`git diff`
    await $`exit 1`
}