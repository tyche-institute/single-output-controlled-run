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

## Results

The results are added to this file after the runs. The records of every run (jobs, logs, releases and provenance
files) are kept in the article's replication package, because GitHub keeps logs for a limited time.

## Licence and origin

Apache License 2.0; see [`LICENSE`](LICENSE) and [`NOTICE`](NOTICE). The jobs of the release workflow and the six
build configurations come from slsa-framework/slsa-verifier at commit `ea584f4502babc6f60d9bc799dbbb13c1caa9ee6`.
Each derived file says so in its header, and `PROTOCOL.md`, Section 3, lists every change.
