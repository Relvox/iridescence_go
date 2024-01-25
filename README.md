# Iridescence (go)

## Content

* `./asserts` - Asserts for `slice` and `map` (with test) and a `testing.T` mock.
  * **Status**: Ready
* `./clients` - `JsonGet` and `JsonPost`.
  * **Status**: Ready
* `./experimental` - Misc stuff that is whatever.
  * **Status**: Stuff...
* `./files` - `path` and `FS` utils.
  * **Status**: Ready
  > + `Split(fullPath) -> dir, name, ext`
  > + `IsolateName("./path/to/file.ext") -> "file"`
  > + `ListFS(...)`
  > + `type SubdirectoryFS ...`
* `./geom` - `Point`s, `Vector`s, and a Staggered `SGrid`.
  * **Status**: Stable but useless.
* `./handlers` - `FSTemplateHandler` magic.
  * **Status**: Needs testing...
* `./logging` - `slog/log` helpers with a `TeeLogger` and `Attr`.
  * Stats: Ready
* `./maths` - `Sign`, `Abs`, `Min`, `Max`, `Sum`, `GeometricMean`, `XenoSum`
  * **Status**: Ready
  * `./maths/test_gen` - Generates the repetitive tests.
* `./middleware` - `LoggingMiddleware`
  * **Status**: Ready
* `./queues` - `THQueue` is an auto-growing Ring Queue.
  * **Status**: Ready
* `./random` - A Source wrapper that add stuff on top of WELL512
  * **Status**: **#TODO:ODOT#**
  * `./random/sources` - Implementation of WELL512 source for `math/rand`.
    * **Status**: Ready
* `./serialization` - `JSON`/`YAML` common SER-DES funcs.
  * **Status**: Ready
* `./servers` - 
  * **Status**: Unholy Mess
* `./sets` - `map[K]struct{}` and related utils. `YAML`/`JSON` and `map`s.
  * **Status**: Ready
* `./templates/funcs` - go `template` helper funcs.
  * **Status**: Ready
* `./text` - `TextGrid` for terminal rendering.
  * **Status**: ?
* `./tools` - Stuff that should probably move
  * `./tools/countula` - a line-of-code counter.
    * **Status**: Ready
  * `./tools/sawmill` - a web based log viewer.
    * **Status**: Needs retesting
* `./utils` - A pile of misc utils.
  * **Status**: I guess...

# TODO

1. [ ] Rearrange to standard layout -> v0.1
2. [ ] Separate out tools and stuff
3. [ ] Monadic Option/Result
   1. [ ] Integrate with stuff