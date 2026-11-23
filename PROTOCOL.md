# Two controlled release runs of a copy of slsa-verifier's release workflow: protocol

**Fixed on 23 November 2026, before the repository of the runs existed and before any run was made.**

Article: "What Exactly Was Verified? Claim-Scoped Verification Coverage in Release Pipelines", by Anton Sokolov,
Tyche Institute, Tallinn, Estonia.

## 1. Why these runs, and what was known before them

The article studies a pattern in GitHub Actions workflows. The legs of a matrix job each write the same output, and a
later job reads that output as one value. Its main real example is the release workflow that accompanies
slsa-verifier v2.7.1: six legs build six binaries, and one later job verifies the binary that the value names.

The article models this release from the workflow's source and shows three links separately:

1. **The output rule.** GitHub documents that the later job receives the value of the last leg to complete
   successfully.
2. **The order.** Which leg completes last varies: in 84 scheduled runs of the same workflow, each of the six legs
   completed last at least 10 times. Scheduled runs skip the verification job.
3. **The check.** The verification job's two shell blocks, run in a model with stand-in commands, check only the
   binary that the value names.

No release run in the article shows the three links together. Its field study found no release run of a workflow
that verifies. The runs that support the output rule are plain matrix jobs, not a matrix of calls to a reusable
workflow, which is the form of this release.

The runs of this protocol close that gap without touching anybody's project: we run a copy of the workflow in a
repository of our own. We contact nobody, and we open no issue, fork or pull request anywhere.

The article reports no vulnerability, unsafe release or signature failure. GitHub documents the rule, the single
check may be meant as a smoke test, and the project's release guide offers maintainers a script that verifies every
file. The question of the article is what a green verification job covers.

**Known before.** We had read the source files (September 2026), the model's results, GitHub's documentation of the
rule, and the completion times of the legs in slsa-verifier's scheduled runs up to 20 November 2026. In those runs the
legs upload nothing and the verification job is skipped. Nothing had been run: when this protocol was fixed, no
repository, tag or workflow run of ours existed. We did not know whether the versions of the builder and the verifier
that the source pins still work today. Section 10 says what happens if they do not.

## 2. Source

- The workflow: `.github/workflows/release.yml` of `slsa-framework/slsa-verifier` at commit
  `ea584f4502babc6f60d9bc799dbbb13c1caa9ee6`; 119 lines; SHA-256
  `1332cd195bb0807f2df9ecee458e8fe5ee9d80609f78f40a9bcef87b9b6ad69a`.
- The six build configurations `.slsa-goreleaser/<os>-<arch>.yml` at the same commit.
- The builder that each leg calls: `slsa-framework/slsa-github-generator/.github/workflows/builder_go_slsa3.yml@v2.0.0`.
- The installer that the verification job uses:
  `slsa-framework/slsa-verifier/actions/installer@3714a2a4684014deb874a0e737dffa0ee02dd647`, which the source
  comments as v2.6.0.

The article's replication package keeps these files with their digests (`workflow-coverage/sources/` and
`workflow-coverage/SOURCE-MANIFEST.json`). All are under the Apache License 2.0.

## 3. The repository and its differences from the source

The repository is `tyche-institute/single-output-controlled-run` on GitHub, public, created for these runs. It holds
a toy Go program without dependencies, six build configurations, two workflows and this protocol.

| File | Relation to the source |
|:--|:--|
| `.github/workflows/release.yml`, first form | The jobs `args`, `builder` and `verification` are lines 22 to 90 of the source, byte for byte. Kept from the rest: the tag trigger, `permissions: read-all` and the `GH_TOKEN` variable. Removed: the manual trigger and the daily scheduled trigger; the two jobs that report the outcome of scheduled runs to the source project's issue tracker, with their two variables. Changed: the workflow's name. Added: a header comment. |
| `.github/workflows/release.yml`, second form | The first form with two changes. A job `plant` is added between `builder` and `verification`. The `needs` of `verification` is `[builder, plant]` instead of `[builder]`, so that the plant comes before the check. The steps of `verification` are unchanged. |
| `.slsa-goreleaser/<os>-<arch>.yml` | The source's six files with three values changed: the prefix of `binary` (`sample` instead of `slsa-verifier`), `dir`, and the variable that `ldflags` sets. Added: a header comment. |
| `.github/workflows/audit.yml` | Not in the source. One job: the installer step of the verification job, then that job's two commands applied to each of the six binaries in turn. It collects all six results and fails if any of them failed. |
| `.github/workflows/toolcheck.yml` | Not in the source. One job: the installer step of the verification job, then a step that prints the verifier's version. It builds and verifies nothing. |
| `cmd/sample/main.go`, `go.mod` | Not in the source: the toy program. |

The six binaries are `sample-linux-amd64`, `sample-linux-arm64`, `sample-windows-amd64.exe`,
`sample-windows-arm64.exe`, `sample-darwin-amd64` and `sample-darwin-arm64`. Each leg's provenance file carries its
binary's name followed by `.intoto.jsonl`.

A script of the article's package, `check_repo.py`, checks every statement of this section against the retained
source files. It ran before the lock (Section 14).

## 4. The runs

A tool check, then two release runs and two audit runs, in this order:

0. **Tool check.** One manual run of `toolcheck.yml` on the main branch. It tests whether the installer that the
   source pins still installs the verifier today. It builds nothing and tests no prediction. If it fails, the
   installer rule of Section 10 applies before the first tag, and the change is reported as a deviation.
1. **Run A, unplanted.** We push tag `v0.1.0` on the commit that holds the first form of the workflow.
2. **Audit A.** After run A has ended, one manual run of `audit.yml` on tag `v0.1.0`.
3. **Run B, planted.** We push tag `v0.2.0` on the next commit, which holds the second form.
4. **Audit B.** After run B has ended, one manual run of `audit.yml` on tag `v0.2.0`.

Run B is made whatever run A shows. Between the two runs only the form of the workflow changes, and what Section 10
allows if a tool fails.

## 5. The plant

In run B, after all six legs have completed, the job `plant` lists the jobs of its own run through GitHub's
application programming interface (API) and computes each leg's completion time (Section 7). It picks the leg that
completed first. Among legs that completed in the same second it picks the one that comes first in the matrix: linux
before windows before darwin, and amd64 before arm64.

It then downloads that leg's binary from the release and records its SHA-256. It replaces the binary in the release,
under the same name, with a marker file of two lines of text. It reads the asset back and compares it with the
marker. It does not touch the provenance files.

The marker's text is fixed, and its SHA-256 is
`03ecfc692bec41015f63d50cf79e947d91f36bd043be5905ce43d58f469923aa` (140 bytes).

The job reads completion times only. It does not read the builder's output, so its choice of the leg cannot depend
on the value under test. The rule picks a leg that did not complete last, unless all six legs complete in the same
second.

The marker stands for any fault that leaves in the release a file that differs from the one its provenance
describes. The verifier must fail for it, since the provenance names another digest. Audit B tests that it does.

## 6. The audit

The audit applies to each of the six binaries the two commands that the verification job applies to one: the download
of the binary and of its provenance file, and `slsa-verifier verify-artifact` with the same source and tag
arguments. It prints each downloaded binary's SHA-256 and one result line per binary, and it fails if any binary
failed. It is the loop form that the article evaluates beside the workflow's single check.

## 7. Definitions

The first five are those of the article's field study (supplemental material, Section S5.4).

- **Leg.** One of the six combinations of the matrix of the job `builder`. A job of the run belongs to the leg
  `<os>-<arch>` when its name begins with `builder-<os>-<arch> / `.
- **Completion time of a leg.** The latest completion time among its jobs, as the API reports it, to the second.
- **Successful leg.** A leg none of whose jobs failed, was cancelled or timed out.
- **Last leg.** The successful leg with the latest completion time. If two or more successful legs share the latest
  second, the run is tied for last.
- **Margin.** The seconds between the completion of the last leg and that of the leg before it.
- **Close run.** A run whose margin is 1 second. The API reports the completion of a job to the second, and the
  service assigns the output a moment after the last job of a leg completes. With 1 second between the two last
  legs we do not treat their order as settled.
- **The value.** `needs.builder.outputs.go-binary-name` as the verification job received it. We read it from that
  job's log: the `ARTIFACT` line in the environment of its steps, and the line "Verifying … using …" that its second
  step prints.
- **Tested run.** A release run is tested for P1 when the verification job started its "Download assets" step: the
  value is then in the log. It is tested for P2 and P3 when the job reached its "Verify assets" step.

## 8. Predictions

- **P1, the output rule (runs A and B).** The value is the binary name of the last leg. In a run tied for last or
  close, the prediction is that the value names one of the legs that completed within 1 second of the latest. Such a
  run is reported as tied or close, and it does not tell those legs apart.
- **P2, the check (runs A and B).** The verification job's log holds exactly one line "Verifying X using
  X.intoto.jsonl", where X is the value. Every line in which the verifier names an artifact names X. The step
  succeeds.
- **P3, fail-open (run B).** The run's conclusion is `success`, although the release held the marker in place of one
  leg's binary before the verification job started. Audit B fails: its result is a failure for the planted binary
  and a pass for each of the other five.
- **Control (audit A).** Audit A succeeds with six passes. The two audits then differ only where the plant is.

Two other readings of the rule are recorded beside P1 for each run:

- the leg whose `build-dry` job completed last, since that job sets the value inside the builder;
- the first and the last leg of the matrix as it is written (linux-amd64 and darwin-arm64).

When such a reading names the same leg as the last-leg rule in a run, that run cannot tell the two apart, and the
report says so.

## 9. What counts against the model

- **Against P1.** In a run that is neither tied for last nor close, the value names a leg other than the last. In a
  tied or close run, the value names a leg that completed more than 1 second before the latest. In any run, the
  value is empty or names no binary of the release. In run B, if the value names the planted leg, the verification
  job meets the marker and the run fails. This also counts against P1, because the plant sits on the first leg to
  complete.
- **Against P2.** More than one "Verifying" line, or an artifact named by the verifier that is not the value.
- **Against P3.** Run B fails at the "Verify assets" step.
- **An ineffective plant.** Run B succeeds and audit B passes for all six binaries. The marker would then not be a
  failing verification, run B would test nothing about P3, and we would report the plant as ineffective.
- **A failed control.** Audit A fails although nothing was planted. The fault is then in the audit or the tools, we
  report it, and the comparison of the two audits loses its control.

## 10. Failures of the tools, and the stopping rule

A run can fail for a reason that says nothing about P1 to P3: a job inside the builder fails, the installer cannot
install the verifier, the toy program does not build, or the `plant` or `audit` job has a defect.

- Every run that we make is reported with its records, whatever it shows.
- A run counts for a prediction only if it is tested for that prediction (Section 7). A run that stopped before
  the verification job's "Download assets" step tests nothing.
- After such a failure we make the smallest change that removes its cause and push the next tag: `v0.1.1`, `v0.1.2`
  and so on for the first form, `v0.2.1` and so on for the second. Each change is reported as a deviation. Allowed
  changes:
  - to the toy program or its configurations;
  - to the `plant` or `audit` job;
  - to the installer's version in `verification` and `audit`: the installer of the latest released slsa-verifier,
    and if that fails as well, a step that builds the verifier from the source of that release;
  - to the builder's version tag: the latest release of the builder.
- No run is repeated because of what it shows about P1, P2 or P3. We do not use GitHub's re-run of failed jobs: a new
  attempt is a new tag.
- At most 6 tags in all. If the limit is reached without two runs tested for P2 and P3, one in each form, we stop
  and report.
- The tool check of Section 4 may be repeated after each change of the installer. It is not a release run and has no
  tag.

## 11. Records

After the runs, `collect.py` reads through GitHub's command-line client:

- each run: identifier, event, tag, commit, attempt, status, conclusion and times;
- its jobs and their steps: names, conclusions, start and completion times;
- the logs of the jobs `verification`, `plant`, `audit` and `toolcheck`, and of each leg's `provenance` and
  `upload-assets` jobs;
- each release with its assets: name, size, digest and times;
- the provenance files of each release, and from each the subject's name and digest;
- the workflow files at each run's commit.

GitHub keeps logs for a limited time. The article's package keeps the copies.

## 12. Analysis

`analyse.py` computes each leg's completion time, the last leg, the margin and ties. It reads the value and the
verifier's lines from the logs, the plant's record and the audits' result lines, and compares them with P1 to P3 and
the control. It writes one result file for people and one for machines.

A self-test (`selftest.py`) ran the analysis on constructed records before the lock: runs that follow the model, a
value that names a leg other than the last, a value that names the planted leg, an empty value, a tie for last, a
close run, a failure before the value is read, a failed control and an ineffective plant. It also ran the shell
blocks of the `plant` and `audit` jobs, as the workflow files hold them, on the job listing of one retained
scheduled run of the source workflow, with stand-ins for the command-line client and the verifier.

## 13. What the runs can and cannot show

These are two runs in one repository, with a toy program, on GitHub as it is on the day. They can show the three
links together in a run, and a file that fails verification under a run that ends in success.

They do not show how often each leg completes last; the scheduled runs of the article do. They do not show what
happened in any release of slsa-verifier, or that any project shipped a bad binary. A copy in our own repository is
not an observation of a project's release, and the article says so wherever it reports these runs.

## 14. Lock, public traces and deviations

**Lock.** Before the repository is created, the article's package records the SHA-256 of this protocol, of every file
of the repository in both forms, and of the scripts and the self-test's log (`controlled-run-2026-11/freezes/C0.json`).
The list is stamped with OpenTimestamps. The repository's first commit holds this protocol, and the tags come after
it.

**Public traces.** The repository, its releases, its runs and their logs are public. For each binary the builder
signs a provenance statement and records it in Sigstore's public transparency log. Such entries are permanent, and
they name the repository, the tag and the commit. The marker stays in the release of run B under the name of a
binary. After the records are collected, the notes of that release and the README name it.

**Deviations.** Every deviation from this protocol is listed, with its time, in the results
(`controlled-run-2026-11/results/RESULTS.md`).
