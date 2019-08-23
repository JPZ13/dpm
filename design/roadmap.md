# Roadmap

- Have volume names be generated and use the first 8 letters of the project path
  hash prefixed by "dpm-"
- Add list command to list whether the project is active and any commands
  defined in the project's dpm file
- Modify the path table to handle blobs as originally intended rather than
  the project info struct
- Add prune command to get rid of the volumes both in a given project folder and
  also have a global option
