# Ent Cyclic Dependency Reproduction

I'm new to Golang, but not new to development, I've been using various languages for about 10 years. I was wondering if I could get some clarification on the Ent Documentation. I am specifically referring to this part of the documentation: https://entgo.io/docs/hooks/#hooks-registration. When I follow the advice in the info box, and follow the steps below it, I can generate my ent hooks just fine. Then when I do `go build` I get a cyclic error:
```
package <project>
        imports <project>/db
        imports <project>/db/runtime
        imports <project>/db/schema
        imports <project>/db: import cycle not allowed
```
Context: I installed ent in a package called `db`, whereas by default the docs use an `ent` folder.

Do I *need* to use the runtime? I know the advice says I **MUST** include it, so am I putting it in the wrong place?

I tried to follow this advice in the Github issues:
- https://github.com/ent/ent/issues/3646#issuecomment-1633376763
- https://github.com/ent/ent/issues/892#issuecomment-1182760541

But I am not comprehending the instructions, so I thought hoping on here would work. I'm learning Golang so some of the vocabulary and necessary context is missing for me to be able to implement what I found in the comments.

I have the runtime imported in a file that exposes a `func ConnectToDb` which is what initializes the Client.

I'm importing it how the docs say with the `import _ "survey/db/runtime"`. I'm assuming I need the runtime, otherwise it wouldn't be generated. I was also reading that an `init` function is called by Golang (during package initialization?), so it feels pretty necessary! lol If I don't include it everything builds.

My project structure is basically like this: `main.go` (entry) -> calls `ConnectToDb` from `survey/db/db.go` which initializes the `db.Client`.

I would appreciate the help! Let me know if I can answer any additional questions.