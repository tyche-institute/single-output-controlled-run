# Controlled release runs of a single-output release workflow

This repository is a research record. It holds a copy of the release workflow of
[slsa-verifier](https://github.com/slsa-framework/slsa-verifier) and a toy program. It exists for two controlled
release runs of that copy and for the audit of each. It ships no software for use.

It belongs to the article "What Exactly Was Verified? Claim-Scoped Verification Coverage in Release Pipelines" by
Anton Sokolov (ORCID [0000-0003-2452-7096](https://orcid.org/0000-0003-2452-7096)), Tyche Institute, Tallinn,
Estonia. A preprint of an earlier version of the article is on Zenodo:
<https://doi.org/10.5281/zenodo.23639536>.

## Before you download anything

One file of the second release is **not a program**. It is a text marker that carries the name of a binary. The job
`plant` put it there on purpose, and [`PROTOCOL.md`](PROTOCOL.md) gives its SHA-256. The other binaries are a toy
program that prints one line. Nothing here is meant to be installed or run.

## What the runs test

In the workflow, six matrix legs build six binaries, and each leg writes the same output. A later job reads that
output as one value and verifies the binary that the value names. The article models this release from the
workflow's source. In the model the value is that of the last leg to complete, so the job verifies one binary of
six, and its green result says nothing about the other five.

The runs test that model on GitHub itself, in this repository only:

1. **Tool check.** `toolcheck.yml` installs the verifier and prints its version.
2. **Run A.** A tag is pushed. The jobs `args`, `builder` and `verification` of `release.yml` are, byte for byte,
   those of the source.
3. **Audit A.** `audit.yml` verifies each of the six binaries of run A's release.
4. **Run B.** The same workflow with one added job, `plant`. After all six legs have completed, it replaces the
   binary of the leg that completed first with the marker. The verification job runs after it, and its steps are
   unchanged.
5. **Audit B.** `audit.yml` verifies each of the six binaries of run B's release.

[`PROTOCOL.md`](PROTOCOL.md) was fixed before this repository existed and before any run. It states the
predictions, what would count against the model, what happens if a tool fails, and what the runs cannot show.

## What this is not

The runs intervene in nobody's project. We opened no issue, fork or pull request anywhere, and we contacted nobody.

The runs do not show what happened in any release of slsa-verifier, or that any project shipped a bad binary. The
article reports no vulnerability, unsafe release or signature failure. GitHub documents how a job reads the output
of a matrix, the single check may be meant as a smoke test, and the project's release guide offers maintainers a
script that verifies every file. The article asks what a green verification job covers.

## Files

| Path | What it is |
|:--|:--|
| `PROTOCOL.md` | The protocol of the runs, fixed before them. |
| `.github/workflows/release.yml` | The copied release workflow. The commit of run A holds its first form, the commit of run B its second form with the job `plant`. |
| `.github/workflows/audit.yml` | The audit: the verification job's two commands, applied to each of the six binaries. |
| `.github/workflows/toolcheck.yml` | The tool check. |
| `.slsa-goreleaser/` | The six build configurations, with three values changed. |
| `cmd/sample/`, `go.mod` | The toy program. |
| `LICENSE`, `NOTICE` | Licence and origin of the copied files. |
| `results/` | The two result files, added after the runs (see Results). |

## Results

The tool check, both release runs and both audits were made on 23 November 2026, in the order of the protocol: two
tags, no repeated run and no re-run of a job.

| | Run A, tag `v0.1.0`, nothing planted | Run B, tag `v0.2.0`, one binary replaced |
|:--|:--|:--|
| Release run | [45236660430](https://github.com/tyche-institute/single-output-controlled-run/actions/runs/45236660430): success | [45237708522](https://github.com/tyche-institute/single-output-controlled-run/actions/runs/45237708522): success |
| Last leg to complete | darwin-amd64, 4 seconds after the leg before it | windows-amd64, 19 seconds after the leg before it |
| Value that the verification job received | `sample-darwin-amd64` | `sample-windows-amd64.exe` |
| Binaries that the verification job checked | that one | that one |
| Planted | nothing | `sample-windows-arm64.exe`, the binary of the first leg to complete, replaced by the marker before the verification job started; the job `plant` completed 4 seconds before it |
| Audit of all six binaries | [45237496257](https://github.com/tyche-institute/single-output-controlled-run/actions/runs/45237496257): success, six passes | [45238563069](https://github.com/tyche-institute/single-output-controlled-run/actions/runs/45238563069): failure, five passes and one failure, for `sample-windows-arm64.exe` |

- **P1, the output rule: held in both runs.** The value named the last leg to complete, a different leg in each run.
- **P2, the check: held in both runs, in the reading of deviation D1 below.** The verification job verified the one
  binary that the value names.
- **P3, fail-open: held in run B.** The run ended in success although the release held the marker in place of one
  binary before the verification job started. The audit failed for exactly that file.
- **Control: held.** The audit of run A's release passed for all six binaries.

**What the two runs do not tell apart.** In both runs the leg that completed last was also the leg whose `build-dry`
job completed last. The runs therefore cannot say which of the two moments decides the value. They do rule out, in
both runs, that the value comes from the first or from the last leg of the matrix as written.

**Deviations from the protocol (2).**

- *D1, the reading of P2.* In the installer step of the verification job, an earlier release of the verifier
  checks the downloaded verifier, and prints a line of the same form as the check of a release binary. The
  protocol's sentence for P2, "Every line in which the verifier names an artifact names X", did not foresee that
  line. By its letter the line counts against P2 in both runs, and the analysis script as locked says so. We report
  P2 for the lines outside the installer step, since P2 is about the binaries of the release, and we keep both
  outputs: [`results/RESULTS.md`](results/RESULTS.md) and
  [`results/locked-script/RESULTS.md`](results/locked-script/RESULTS.md). We wrote the changed reading after run A
  had ended and before run B's verification job started.
- *D2, the collection.* The script that reads the records stopped twice because GitHub answered one request for a
  log with HTTP 503. A wrapper that repeats such a read collected the records in one pass.

**The files of the releases.** In the release of `v0.2.0`, `sample-windows-arm64.exe` is the marker: 140 bytes of
text, SHA-256 `03ecfc692bec41015f63d50cf79e947d91f36bd043be5905ce43d58f469923aa`. It is not a program, and its
provenance file describes the binary that the run built, so its verification fails. Every other binary of both
releases is the toy program and verifies against its provenance.

**Where the records are.** The tool check is run [45236053013](https://github.com/tyche-institute/single-output-controlled-run/actions/runs/45236053013). The two result files in `results/`
are as the analysis wrote them; they name scripts and folders of the article's replication package, which keeps the
scripts and the records of every run (jobs, logs, releases and provenance files), because GitHub keeps logs for a
limited time. The runs say nothing about any release of slsa-verifier: they are runs of a copy, with a toy program,
in this repository.

## Licence and origin

Apache License 2.0; see [`LICENSE`](LICENSE) and [`NOTICE`](NOTICE). The jobs of the release workflow and the six
build configurations come from slsa-framework/slsa-verifier at commit `ea584f4502babc6f60d9bc799dbbb13c1caa9ee6`.
Each derived file says so in its header, and `PROTOCOL.md`, Section 3, lists every change.
