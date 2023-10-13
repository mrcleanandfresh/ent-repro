# Ent Cyclic Dependency Reproduction

I'm new to Golang, but not new to development, I've been using various languages for about 10 years. 

I could use some clarification on the Ent Documentation. I am specifically referring to this part of the documentation: <https://entgo.io/docs/hooks/#hooks-registration>. When I follow the advice in the info box, and follow the steps below it, I can generate my ent hooks just fine without a cyclic error like the docs mention. Then when I do `go build` I get a cyclic error:
```
package test
        imports test/db
        imports test/db/runtime
        imports test/db/schema
        imports test/db: import cycle not allowed
```
Context: I installed ent in a package called `db`, whereas by default the docs use an `ent` folder.

Do I *need* to use the runtime? I know the advice says I **MUST** include it, so am I putting it in the wrong place? Or misconfiguring something else?

I tried to follow this advice in the Github issues:
- https://github.com/ent/ent/issues/3646#issuecomment-1633376763
- https://github.com/ent/ent/issues/892#issuecomment-1182760541

But I am not comprehending the instructions, so I thought posting the minimum reproduction for Discord would work. I'm learning Golang so some of the vocabulary and necessary context is missing for me to be able to implement what I found in the comments.

Please take a look at the code and let me know in Discord what I'm doing incorrectly! Thank you ahead of time.
