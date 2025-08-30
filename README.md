# Simple Typeahead Suggestions (using Trie)

The simplest, non-scalable feature (implemented using Trie data structure), that suggests next words of queries based on user input.

## This design is extremely inefficient

**Yes**.

From System-Design perspective, this can be very intersting design homework. However, this current project only focuses on using basics of Trie functionality to implement the simplest possible typeahead sugestion feature. Hence this implementation has following flaws, compared to a full-fledged typeahead / autocomplete feature used in search engines like Google, Bing etc.

- Non-scalability
- No suggestions based on geographical location of user
- No suggestions based on current trends
- No spell-check
- Support for only English language

### ERROR IN IMPLEMENTATION LOGIC

I have made a logical error in the dictionary heap implemention during initial insertion
