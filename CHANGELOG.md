# StringSageMap Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [v0.1.0] - 2019-12-27
### Added
- `SafeStringMap` structure. Defines a map of strings that are managed by a `sync.RWLock` to ensure safe Read-Write access in concurrent environments. All read operations utilse `RLock`. All write operations utilise `Lock`.
- `NewSafeStringMap` function creates a new instance of a `SafeStringMap` structure.
- `SafeStringMap.Load` returns the string value associated with the requested key if it exists in the map. If not an emptry string is returned and the functions `bool` result will be set to `false`.
- `SafeStringMap.Delete` removes a value from the map that is associated with the specified key value - if it exists.
- `SafeStringMap.Store` adds a new key value pair to the map if it doesnt already exist. If the key already exists, the provided string value will be used to update (overwrite) the value associated with the key.
- `SafeStringMap.Count` returns the count of strings stored within the map (as an integer)
### Changed
### Deprecated
### Removed
### Fixed
### Security

[v0.1.0]: 
https://github.com/BigOokie/safestringmap/releases/tag/v0.1.0
