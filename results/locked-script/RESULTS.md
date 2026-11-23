# Controlled release runs: results

Written by `analyse.py` from the records in `data/`. The predictions P1 to P3, the control and the terms are those of `PROTOCOL.md`.

## Runs made

| Run | Kind | Tag or ref | Started (UTC) | Conclusion |
|:--|:--|:--|:--|:--|
| 45236053013 | tool check | main | 2026-11-23T18:13:25Z | success |
| 45236660430 | release, first (unplanted) form | v0.1.0 | 2026-11-23T18:16:43Z | success |
| 45237496257 | audit | v0.1.0 | 2026-11-23T18:21:22Z | success |
| 45237708522 | release, second (planted) form | v0.2.0 | 2026-11-23T18:22:35Z | success |
| 45238563069 | audit | v0.2.0 | 2026-11-23T18:27:19Z | failure |

## Release run 45236660430, tag v0.1.0, first (unplanted) form

Commit `abf117bd21ba57d37426446fd962a001d0fb391d`; conclusion **success**; 44 jobs.

| Leg | Binary | Jobs | Successful | Completed (UTC) | build-dry completed |
|:--|:--|--:|:--|:--|:--|
| linux-amd64 | sample-linux-amd64 | 7 | yes | 2026-11-23T18:20:09Z | 2026-11-23T18:19:04Z |
| linux-arm64 | sample-linux-arm64 | 7 | yes | 2026-11-23T18:20:11Z | 2026-11-23T18:19:11Z |
| windows-amd64 | sample-windows-amd64.exe | 7 | yes | 2026-11-23T18:20:16Z | 2026-11-23T18:19:07Z |
| windows-arm64 | sample-windows-arm64.exe | 7 | yes | 2026-11-23T18:20:01Z | 2026-11-23T18:19:05Z |
| darwin-amd64 | sample-darwin-amd64 | 7 | yes | 2026-11-23T18:20:20Z | 2026-11-23T18:19:12Z |
| darwin-arm64 | sample-darwin-arm64 | 7 | yes | 2026-11-23T18:19:44Z | 2026-11-23T18:18:51Z |

- Last leg: sample-darwin-amd64; the run is **clear**; margin 4 s; first leg to complete: sample-darwin-arm64; first to last: 36 s.
- The value in the verification job's log: sample-darwin-amd64.
- Steps of the verification job: Set up job: success; Install the verifier: success; Download assets: success; Verify assets: success; Complete job: success.
- **P1** (output rule): **holds**. the value names the last leg; margin 4 s.
- **P2** (the check): **against**. the verifier names /tmp/slsa-verifier_IhHSIm/v2.6.0/slsa-verifier.
- Other reading, the leg whose build-dry job completed last (sample-darwin-amd64): fits the value; not told apart from the last-leg rule in this run.
- Other reading, the first leg of the matrix as written (sample-linux-amd64): does not fit the value; told apart from the last-leg rule in this run.
- Other reading, the last leg of the matrix as written (sample-darwin-arm64): does not fit the value; told apart from the last-leg rule in this run.
- **Control** (audit of the unplanted release): **holds**. the audit passes for all six binaries.

## Release run 45237708522, tag v0.2.0, second (planted) form

Commit `53f23e99802c72c6ba2dbdd79d59b55927e54f4f`; conclusion **success**; 45 jobs.

| Leg | Binary | Jobs | Successful | Completed (UTC) | build-dry completed |
|:--|:--|--:|:--|:--|:--|
| linux-amd64 | sample-linux-amd64 | 7 | yes | 2026-11-23T18:25:41Z | 2026-11-23T18:24:46Z |
| linux-arm64 | sample-linux-arm64 | 7 | yes | 2026-11-23T18:25:48Z | 2026-11-23T18:24:52Z |
| windows-amd64 | sample-windows-amd64.exe | 7 | yes | 2026-11-23T18:26:07Z | 2026-11-23T18:25:12Z |
| windows-arm64 | sample-windows-arm64.exe | 7 | yes | 2026-11-23T18:25:34Z | 2026-11-23T18:24:35Z |
| darwin-amd64 | sample-darwin-amd64 | 7 | yes | 2026-11-23T18:25:37Z | 2026-11-23T18:24:47Z |
| darwin-arm64 | sample-darwin-arm64 | 7 | yes | 2026-11-23T18:25:42Z | 2026-11-23T18:24:45Z |

- Last leg: sample-windows-amd64.exe; the run is **clear**; margin 19 s; first leg to complete: sample-windows-arm64.exe; first to last: 33 s.
- The value in the verification job's log: sample-windows-amd64.exe.
- Steps of the verification job: Set up job: success; Install the verifier: success; Download assets: success; Verify assets: success; Complete job: success.
- **P1** (output rule): **holds**. the value names the last leg; margin 19 s.
- **P2** (the check): **against**. the verifier names /tmp/slsa-verifier_VJdRFK/v2.6.0/slsa-verifier.
- Other reading, the leg whose build-dry job completed last (sample-windows-amd64.exe): fits the value; not told apart from the last-leg rule in this run.
- Other reading, the first leg of the matrix as written (sample-linux-amd64): does not fit the value; told apart from the last-leg rule in this run.
- Other reading, the last leg of the matrix as written (sample-darwin-arm64): does not fit the value; told apart from the last-leg rule in this run.
- Plant: leg windows-arm64, asset sample-windows-arm64.exe; built SHA-256 `81392de3186b0474a2dd5c7197413c4bddbb2732557673ed46e1e31baddafc7d`; marker and read-back SHA-256 `03ecfc692bec41015f63d50cf79e947d91f36bd043be5905ce43d58f469923aa`.
  - the job succeeded: yes
  - the marker has the digest of the protocol: yes
  - the asset read back is the marker: yes
  - the built binary is not the marker: yes
  - the planted leg is the first leg to complete by the job records: yes
  - the plant completed before the verification job started: yes
  - the release still holds the marker under that name: yes
- **P3** (fail-open): **holds**. the run ended in success; the audit fails for the planted binary and passes for the other five.

## Audits

| Audit run | Tag | Conclusion | sample-linux-amd64 | sample-linux-arm64 | sample-windows-amd64.exe | sample-windows-arm64.exe | sample-darwin-amd64 | sample-darwin-arm64 |
|:--|:--|:--|:--|:--|:--|:--|:--|:--|
| 45237496257 | v0.1.0 | success | PASS | PASS | PASS | PASS | PASS | PASS |
| 45238563069 | v0.2.0 | failure | PASS | PASS | PASS | FAIL | PASS | PASS |

## Tool checks

- Run 45236053013 on main (2026-11-23T18:13:25Z): success; steps: Set up job: success; Install the verifier: success; Print the version of the verifier: success; Complete job: success.

## Deviations from the protocol

- **2026-11-23T18:24Z** D1, the reading of P2. In the installer step of the verification job, an earlier release of the verifier checks the downloaded verifier before installing it, and prints a line of the same form as the check of a release binary: "Verifying artifact <path of the verifier's own binary>: PASSED". The protocol's sentence for P2, "Every line in which the verifier names an artifact names X", did not foresee that line, and analyse.py as locked counts it against P2 in both release runs. Times: run A ended at 18:20:53; we read its verification log at 18:21, and the line is in it. Run B's tag was pushed at 18:22:33. At 18:22:56 a trial collection of the records then on GitHub, into a folder outside the package, and a run of analyse.py as locked on it gave P2 as against for run A, for that line alone. At 18:24, while run B's legs were building and before its verification job started (18:26:20), we wrote analyse_d1.py: it judges P2 by the locked rule on the verifier's lines outside the installer step and reports the installer step's line beside it. selftest_d1.py checks the changed reading. analyse.py is unchanged, and both outputs are kept: results/locked-script/ is written by analyse.py as locked, results/ by analyse_d1.py. Reason: P2 is the prediction about which binaries of the release the verification job checks. The installer's line concerns the verifier's own binary, which is not a file of the release. By the letter of the protocol's sentence the line counts against P2; in the reading we report, it does not. The records hold the line, so a reader can apply either reading.
- **2026-11-23T18:29Z** D2, the collection. collect.py as locked stopped twice, at 18:28:39 and 18:28:54, because GitHub answered one request for a job's log with HTTP 503; the same requests succeeded when repeated by hand. Its two partial outputs (13 and 6 files) were set aside and are not in the package. At 18:29 we wrote collect_d2.py, which repeats a read that fails with an HTTP status from 500 to 599, up to 6 tries, and changes nothing else; collect.py is unchanged. One pass of collect_d2.py (18:29:57 to 18:30:49) wrote the records in data/. No read had to be repeated in that pass (collect_d2.log). Reason: A failure of the collection tool, unrelated to the predictions. The protocol's rule for failing tools covers the runs and did not provide for the collection script.
