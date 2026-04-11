# The Dothog Manifesto

**Being the Collected Sacred Texts of the Hypermedia Novices, Hidden Disciples of the Honorable ROY T. FIELDING (Whose Dissertation We Have Read, Unlike You)**

---

**_ALL STATEMENTS ARE TRUE IN SOME SENSE, FALSE IN SOME SENSE, MEANINGLESS IN SOME SENSE, TRUE AND FALSE IN SOME SENSE, TRUE AND MEANINGLESS IN SOME SENSE, FALSE AND MEANINGLESS IN SOME SENSE, AND TRUE AND FALSE AND MEANINGLESS IN SOME SENSE. EXCEPT FOR THIS ONE, WHICH IS A MANIFESTO._**

---

## Preamble

We, the Novices, hold these truths to be self-evident: that the `<a>` tag was endowed by its creator with certain unalienable affordances; that among these are linking, navigation, and the pursuit of HATEOAS; that whenever any framework becomes destructive of these ends, it is the right of the developer to abolish it and return to HTML.

We did not set out to write a manifesto. A manifesto implies a movement, and a movement implies organization, and organization implies a Slack channel, and a Slack channel implies someone asking "has anyone tried the new React Server Components?" and at that point we have already failed.

This is not a movement. This is a collection of observations about how the web works, how it was designed to work, and the extraordinary lengths to which the industry has gone to make it work some other way.

We observe. We return HTML. We do not judge.

_(We judge a little.)_

Within these pages, two spirits contend — as they have contended since the first developer mistook a pattern catalog for a building code. In the margins of The Dissertation, in a hand that is not Fielding's and a font that does not exist, their dialogues were found. They are called **The Geodesist** and **The Pattern Master**: one who builds structures that enclose the most space with the least material, and one who builds structures that enclose the least space with the most material, _but very cleanly_. Their encounters are recorded here alongside the other sacred texts, not because they are canonical, but because they keep happening.

---

## I. THE PENTAVERB, Or The Five Commandments of Hypermedia

_The PENTAVERB was discovered by the hermit developer Zarathud-2 in the Fifth Year of The Framework Churn. He found them encoded in a `robots.txt` file while refactoring his cave. Their import was lost, for they were written in `application/xhtml+xml`. However, after 10 weeks and 11 hours of intensive scrutiny, he discerned that the message could be read by any standard-compliant user agent._

> **I** -- There is no Conditions of the PENTAVERB that permits a Novice to build a Single-Page Application. However, comma, if a Single-Page Application is what the problem GENUINELY requires (collaborative editors, design tools, canvas rendering, and other forms of client-side purgatory), then it is permitted, and the Novice shall not feel guilty, for guilt is a client-side state and the server does not care about client-side state. The server has never cared about client-side state. This is the First Lesson and also the Last Lesson and also most of the lessons in between.
>
> **II** -- A Novice shall Partake Joyously of the `<a>` tag; this Devotional Practice to Remonstrate against the popular Paganisms of the Day: of Reactism (no links, only `onClick`), of Angularity (no links, only `routerLink`), of the Vuedoo Priests (no links, only `@click`), and of the Novices themselves (no Dot Hog Buns). _A Novice shall Partake of No Dot Hog Buns, for Such was the Solace of Our Architect when He was Confronted with The Original Snub (see: JSON-over-HTTP being called "REST"). It is not known why Dot Hog Buns were singled out. It is not our place to ask. Some constraints are architectural. Some are dietary. The wise Novice does not distinguish between them._
>
> **III** -- A Novice is Required to keep an `Accept: text/html` header in their heart and a copy of [Chapter 5](https://roy.gbiv.com/pubs/dissertation/fielding_dissertation.pdf) on their desk, which they shall refer to as "The Text" in casual conversation and "The Dissertation" in formal settings and "That Thing I Keep Telling You To Read, Kevin" in code reviews. A Novice who has not read The Dissertation is still a Novice. A Novice who has read The Dissertation is also still a Novice. Reading The Dissertation does not change your rank because THERE IS ONLY ONE RANK. You are already it. Congratulations. Please stop asking about promotions.
>
> **IV** -- A Novice is Prohibited from Believing What They Read in blog posts about "RESTful APIs" that describe JSON endpoints with no hypermedia controls. Whose API is REST? Where is its HATEOAS? WHERE ARE THE LINKS, KEVIN? A Novice who encounters such a blog post is advised to close the tab, take a walk, and remember that the `<a>` tag has been linking documents since before the author of that blog post was born. If the author is older than the `<a>` tag, they have even less of an excuse.
>
> **V** -- A Novice is Required to think for themselves, despite -- or perhaps because of -- all the conditions outlined above. If the PENTAVERB seems to contradict itself, you have begun to understand the PENTAVERB. If the PENTAVERB seems perfectly consistent, read it again, because you missed something. If after reading it five times it STILL seems consistent, you may be experiencing enlightenment. Or a caching issue. These are indistinguishable from the outside.

_The Geodesist, upon reading the PENTAVERB, said: "Five is correct. Five struts meet at every vertex of an icosahedron, and the icosahedron is the most efficient structure in nature. Five commandments, like five edges, enclose maximum wisdom with minimum material." The Pattern Master, upon reading the PENTAVERB, said: "Five is insufficient. I count at least seventeen cross-cutting concerns unaddressed." He was still counting when The Geodesist had finished building._

---

## II. The Novice's Creed

_To be recited before each sprint planning, or silently, while staring at a `package-lock.json`._

> We believe in one Server, the Almighty,
> source of all state, of all things visible and rendered.
>
> We believe in one Protocol, HTTP,
> the only-begotten means of transfer,
> born of the IETF before all frameworks,
> Light of Light, true REST of true REST,
> specified, not invented, one in Being with the Web.
> Through it all requests were made.
> For us and for our users it came down from the server:
> by the power of HTMX it was made incarnate in HTML,
> and became hypertext.
>
> For our sake it was misnamed under Conditions of JSON;
> it suffered misuse and was buried under `node_modules`.
> On the third RFC it rose again
> in accordance with the Dissertation;
> it ascended into the browser
> and is seated at the right hand of the `<body>`.
> It will come again in `200 OK` to render the living and the cached,
> and its responses will have no end.
>
> We believe in the Hypermedia, the Link, the giver of state,
> who proceeds from the Server and the Representation,
> who with the `GET` and the `POST` together is worshipped and glorified,
> who has spoken through the `<a>` tag.
> We believe in one, holy, uniform, and RESTful Interface.
> We acknowledge one `Content-Type` for the rendering of documents.
> We look for the resurrection of the `<form>`,
> and the life of the web to come.
>
> Amen. Or rather, `200 OK`.

_The Creed is not mandatory. A Novice is Prohibited from Believing What They Read, including this Creed, including this sentence, including the prohibition itself. If you find this paradox troubling, you are not yet ready. If you find it funny, you are also not yet ready, but at least you're having a good time._

---

## III. The Wisdom of the Uniform Interface

_Being a REVELATORY CATECHISM of the Dissertation, in which a Conditions of Fnord Seeker (hereafter "The Fool") poses questions to THE TEXT and THE TEXT answers, sometimes patiently, sometimes not, and once with what can only be described as divine sarcasm._

_These dialogues were recovered from a corrupted `.tar.gz` on a decommissioned CERN proxy server in 2003. The file was labeled `DO_NOT_READ.html`. It contained no JavaScript. It was, and remains, the most well-structured document on the machine. Its MIME type was `text/html`. Its `Content-Type` was `you-are-not-ready`. The IETF has no comment. The W3C has seventeen comments, all contradictory. Tim Berners-Lee was reached for verification and replied only "yes, that sounds about right."_

_Their authenticity is disputed. Their contents are not. The reader is advised to keep a glass of water nearby, because the truth is very dry._

---

**THE FOOL asked: "What is a resource?"**

HEAR ME. A resource is not a row in your database. A resource is not a file on your disk. A resource is not an endpoint. A resource is not a model. A resource is any concept important enough to be named, and naming is the first and most dangerous act of computing (cf. DNS, variable naming, and the entire history of identity management).

A to-do item is a resource. The list of to-do items is a resource. The act of _searching_ for to-do items is a resource. Your anxiety about to-do items is not a resource, but only because you have not yet given it a URI. If you give it a URI, it becomes one. This is the terrible power of naming. The Buddhists tried to warn you about attachment to names. You didn't listen. You were too busy naming your microservices.

_(Editor's note: the original manuscript has "naming things is one of the two hard problems" in the margin, crossed out, with "three hard problems" written above it, crossed out, with "five hard problems" written above THAT, in increasingly frantic handwriting.)_

---

**THE FOOL asked: "What is a representation?"**

A resource is the thing-in-itself, the Ding an sich, the Platonic form of your user table. You never touch it directly. You cannot. You interact with it only through _representations_, which are how the thing chooses to speak to you.

The same resource may speak HTML to a browser and JSON to an API client and plain text to `curl` at 2 AM when you are debugging something that worked in staging. The representation is not the resource any more than a photograph is the mountain. BUT -- and here is where it gets interesting, and by "interesting" I mean "the part everyone skips" -- unlike a photograph, a representation carries CONTROLS. Instructions. _Affordances._ The next available actions. A photograph of a mountain does not contain a trail map. A proper representation does. Your JSON endpoint is a photograph of a mountain thrown at the client's face with a note that says "figure out the trails yourself, the docs are on Confluence, I think, ask Kevin."

---

**THE FOOL asked: "What is the uniform interface?"**

I WILL TELL YOU A SECRET THAT IS NOT A SECRET BECAUSE IT HAS BEEN IN PLAIN SIGHT SINCE 1991 AND YET SOMEHOW THE ENTIRE INDUSTRY HAS MANAGED TO NOT NOTICE IT:

How many web browsers know the difference between a banking application and a wiki?

_None of them._

NONE. Not one. Zero. And yet -- AND YET -- they operate both. They operate ALL OF THEM. EVERY WEBSITE THAT HAS EVER EXISTED. Your browser does not download a BankingApplicationSDK. Your browser does not read the WikiClientDocumentation. Your browser speaks HTTP, understands media types, and follows links. Three things. That is the uniform interface. It works on every website that has ever existed or ever will exist and it has been working this way for THIRTY YEARS and you -- _you specifically, reading this_ -- are building something WORSE than this ON PURPOSE because someone on Medium told you to.

_(The preceding paragraph was found carved into a bathroom stall at an IEEE conference. Below it, in different handwriting: "but what about GraphQL?" Below that, in a third hand: "WHAT ABOUT IT.")_

---

**THE FOOL asked: "What is hypertext?"**

Hypertext is the simultaneous presentation of information and controls such that the information BECOMES THE AFFORDANCE through which choices are obtained and actions are selected.

Read that again. No, actually read it. I know you skimmed it. Everyone skims it. This is why we are in this mess.

Your JSON endpoint returns data. Your HTML page returns data AND WHAT TO DO WITH IT. The `<a>` tag says "here is somewhere you can go." The `<form>` tag says "here is something you can submit." The `<button>` with an `hx-delete` says "here is something you can destroy, and here is exactly how to destroy it, and here is where the confirmation dialog will appear, and none of this required a README."

The client does not need a separate document explaining the API because THE API IS EXPLAINING ITSELF RIGHT NOW, IN THE HTML, WHICH YOU ARE NOT READING BECAUSE YOU ARE WRITING A SWAGGER SPEC. Kevin. Kevin, put down the Swagger spec.

---

**THE FOOL asked: "What is HATEOAS?"**

Hypermedia As The Engine Of Application State.

Yes, it is an ugly acronym. The truth is not always beautiful. Sometimes the truth is an ugly acronym that you should have tattooed on the inside of your eyelids.

The server sends a representation. The representation contains links and forms. The client follows them. THAT IS THE ENTIRE INTERACTION MODEL. I will not be taking questions at this time because I just answered all of them.

The client does not know URLs in advance. The client does not construct URLs from templates it found in the docs. The client does not read your Swagger documentation and hardcode endpoints into a generated API client that you then version and distribute and maintain and deprecate and that breaks every time you rename a field. The client receives a page, sees what it can do, and does it. Like a person. Using a website. Remember? Remember websites? You used one today. It worked. Because of this. BECAUSE OF EXACTLY THIS.

_(fn. 23: It has been observed that the word "HATEOAS" looks like it should be pronounced "hate-ee-ohs," like a breakfast cereal. This is appropriate. It should be part of a balanced architecture.)_

---

**THE FOOL asked: "What does it mean to be stateless?"**

Each request from client to server must contain ALL of the information necessary to understand the request. The server does not remember you. The server does not pine for you between requests. The server is not sitting there holding your session like a loyal retriever waiting by the door. The server has already forgotten you. The server has _moved on._

"But what about sessions?" you cry into the void.

Sessions are the server remembering you. This is a violation. It is a warm, comfortable, widely-practiced violation, like jaywalking or using `!important` in CSS. Everyone does it. The constraint still exists. The question is not whether you violate it -- because you will, you WILL, I have seen the future and it contains `express-session` -- the question is whether you understand what you lost when you reached for it. What you lost is the ability to route any request to any server. What you lost is the ability to cache freely. What you gained is a 32-byte cookie and the illusion of continuity. Whether this trade was worth it is between you and your load balancer.

---

**THE FOOL asked: "What is a media type?"**

A media type is a COVENANT. A sacred compact. A pinky promise between systems. It says: "if you receive `text/html`, here is how you shall process it." It defines the processing model, the structure, the controls, the semantics. It is _the entire instruction set_ for how to handle what you received.

Now. When your API returns `application/json`, what does that tell the client about what to do next?

NOTHING.

_Absolutely nothing._

JSON is a serialization format. It has no links. It has no forms. It has no inherent controls. It is a way of arranging curly braces. You have returned data in a format that carries NO INSTRUCTIONS and then written a SEPARATE DOCUMENT explaining the instructions and then put that document on a SEPARATE WEBSITE and then wondered why your clients BREAK when you CHANGE THE INSTRUCTIONS. You have invented a problem. You have then sold yourself the solution. And the solution is more JSON. It's JSON all the way down. You are in the JSON hole and you are digging.

_(There is a school of thought that `application/hal+json` and similar hypermedia JSON formats solve this problem. There is a school of thought that the `<a>` tag solved this problem in 1993. These schools do not talk to each other. One has a lot more students. The other has a lot less `node_modules`.)_

---

**THE FOOL asked: "What is out-of-band information?"**

Out-of-band information is THE CONSPIRACY. It is the hidden knowledge. The secret handshake. The unspoken assumption.

If your client must read your API docs to know which URL to `POST` to, that is out-of-band. If your client must know that `/api/v2/users/{id}` is the pattern for user resources, that is out-of-band. If your client must be recompiled when you rename a resource, you have coupled the client to the server's URI structure _and you will maintain this coupling in blood and tears until one of you is decommissioned_. This is not a metaphor. I have seen the Jira tickets. They do not end.

The whole point -- the ENTIRE POINT -- of hypermedia is that the server tells the client what to do next IN THE RESPONSE ITSELF. The links are RIGHT THERE. In the HTML. They have been there this whole time. You have been stepping over them to get to your OpenAPI generator. You are the man in the flood who refuses the boat, the helicopter, and the `<a href>`, waiting for God to send you a properly versioned REST client with TypeScript bindings.

---

**THE FOOL asked: "But everyone calls their JSON API 'RESTful.' Are they wrong?"**

If the engine of application state is not being driven by hypertext, then it cannot be RESTful and cannot be a REST API.

Period.

Full stop.

End of transmission.

`Connection: close`

What they have built is RPC. With nice URLs. Some of them have built RPC with nice URLs and an OpenAPI spec and a code generator and a client library and a versioning scheme and a deprecation policy and a migration guide and a breaking changes newsletter, which is to say they have constructed an ENORMOUS and MAGNIFICENT cathedral of infrastructure for the sole purpose of avoiding putting `<a href>` in a response. The engineering effort is genuinely impressive. It is also genuinely unnecessary. It is the Rube Goldberg machine of distributed systems. The ball rolls down the chute, hits the domino, rings the bell, feeds the hamster, and the hamster _renders a table of users._

_(Certain agencies within the United States government have been observed to use the term "RESTful" to describe SOAP endpoints with JSON payloads routed through an API gateway with OAuth2 and rate limiting. We do not name these agencies. They know who they are. We pray for their developers.)_

---

**THE FOOL asked: "Why does this matter? My JSON API works fine."**

"Works fine." "Works fine," they say. Your Titanic also "works fine" right up until the moment it doesn't, and then you're standing on the deck watching your client integrations slide into the North Atlantic because you renamed a field from `userName` to `username` and seventeen services are down and Kevin is on PTO.

REST is software design on the scale of decades. DECADES. Every constraint is intended to promote longevity and independent evolution. Many of the constraints directly oppose short-term efficiency. This is by design. This is ON PURPOSE. If you are building something that will last six months, do whatever you want. Use smoke signals. Store state in a CSV. Put the database password in the URL. I do not care. You will not be around long enough for it to matter.

But if you are building something that must _evolve_ -- while clients depend on it, while teams change, while requirements shift, while Kevin goes on PTO and comes back and the new Kevin doesn't know the old Kevin's conventions -- then you need an architecture that permits change without breaking the contract. Hypertext IS that contract. The representation tells the client what is possible RIGHT NOW. When what is possible changes, the representation changes, and the client adapts, because the client was never hardcoded to anything except "follow the links."

This is not theory. This is how the Web has worked for thirty years. You are standing on the miracle and complaining that the ground is too stable.

---

**THE FOOL asked: "What should I do?"**

Enter the application with a single URI and a set of standardized media types. Follow the links. Submit the forms. Let the server drive the state. That is all.

**THE FOOL said: "That seems too simple."**

Yes. That is the point. That has always been the point. The Dissertation is 180 pages long not because the idea is complicated but because Fielding had to PROVE it was simple, to an industry that has a financial and psychological incentive to believe that things must be complicated, because if things are simple then what have we all been doing for the last twenty years?

_(Do not answer that question. The answer is "inventing problems." You know this. I know this. The `node_modules` directory knows this. It is 900 megabytes of knowledge.)_

---

_THE FOOL closed the PDF. THE FOOL opened their editor. THE FOOL deleted the API client. THE FOOL deleted the route constants. THE FOOL deleted the URL builder. THE FOOL deleted the TypeScript interfaces. THE FOOL wrote an `<a>` tag._

_The browser followed it._

_It worked._

_It had always worked._

_THE FOOL was no longer THE FOOL. THE FOOL was now a web developer. An actual one. The first one in years._

_Somewhere, a `node_modules` directory shrank by 900 megabytes. Nobody noticed, because noticing would require client-side state management, and there wasn't any._

_It is recorded in the margin that The Geodesist was watching from a nearby terminal. He said: "The Fool deleted more than code. Every deletion was an act of ephemeralization — doing more with less until you do everything with nothing. The application now does more than it did before, not despite having less code, but because of it." The Pattern Master was also watching. He opened a pull request reverting the deletions, titled "Restore Separation of Concerns and Reintroduce Repository Layer." It was not merged. It has 47 comments. It is still open._

---

## IV. The Encounters of the Geodesist and the Pattern Master

_Being the Dialogues, Parables, and Confrontations of two figures found arguing in the margins of The Dissertation, in the tradition of the ZHUANGZI, in which **The Geodesist** (who builds domes) encounters **The Pattern Master** (who builds abstractions), and in which it is revealed that one of them has been doing more with less while the other has been doing less with more._

_The Geodesist is known by his proverb: "I seem to be a verb." The Pattern Master is known by his: "See also: Chapter 7." The Geodesist was once asked what he would put on his tombstone. He said: "CALL ME TRIMTAB." The Pattern Master was once asked what he would put on his tombstone. He said: "SEE ALSO: Appendix B, wherein the full pattern language is enumerated with cross-references to the relevant Gang of Four entries." Both understood architecture. One of them understood `<a href>`._

_Their dialogues are presented here without commentary, for the same reason the Zhuangzi is presented without commentary: the stories are the commentary. If you require a commentary on the commentary, you may be The Pattern Master._

---

**THE GEODESIST dreamed he was a Single-Page Application.**

He was a happy SPA, fluttering about, managing his own state, routing his own URLs, rendering himself from a virtual DOM. He did not know he was The Geodesist. Suddenly he woke up, and there he was, solid and unmistakable — a server returning HTML in a single binary.

Now he did not know: was he The Geodesist who had dreamed he was a SPA, or was he a SPA dreaming it was a server returning HTML?

The Pattern Master, hearing this, said: "This is the transformation of concerns. You should implement the Strangler Fig Pattern to progressively migrate from dream state to waking state, maintaining a facade layer that routes requests to whichever consciousness is currently deployed."

The Geodesist said: "I was telling you about the nature of identity and you are selling me a migration strategy."

The Pattern Master said: "All problems are migration problems, properly understood."

The Geodesist said nothing for a long time.

Then The Geodesist said: "That is the most enterprise sentence ever spoken. You have taken the butterfly dream — the most beautiful question in philosophy — and turned it into a JIRA ticket."

_(The Geodesist's tombstone reads "CALL ME TRIMTAB." The Pattern Master's tombstone reads "SEE ALSO: Chapter 7." Neither is wrong. One is a single binary.)_

---

**THE GEODESIST was rendering HTML.**

His template execution followed the natural structure of the response — `<header>`, `<main>`, `<footer>` — never forcing, never fighting the grain of the DOM. He had not replaced his template function in nineteen years.

The Pattern Master watched and said: "Impressive. But you should extract a `TemplateRendererInterface`, implement it with a `ConcreteHTMLTemplateRenderer`, configure it through a `RendererFactory`, and inject it via a `DependencyContainer` so you can swap rendering strategies at runtime."

The Geodesist did not look up. He continued rendering.

The Pattern Master said: "What if you need to render to PDF?"

The Geodesist said: "Do I need to render to PDF?"

The Pattern Master said: "You might."

The Geodesist said: "Cook Ding the butcher carved ten thousand oxen with the same knife. His knife is as sharp as the day it was forged, because he cuts along the natural grain and never forces the blade against bone. You are asking me to carry nineteen knives — one for each animal I have never carved and may never carve — and a `KnifeFactory` to select between them, and a `KnifeStrategyProvider` to configure the factory. The weight of all these knives I do not need has dulled the one knife I do."

The Pattern Master said: "But what about the Open/Closed Principle?"

The Geodesist said: "My template is open to data and closed to your anxiety about data."

---

**THE GEODESIST and The Pattern Master walked through a forest of HTML elements.**

They passed `<div>` after `<div>`, each laden with `onClick` handlers, each heavy with `data-` attributes, each groaning under the weight of its sixteen-word `className`.

Then they came upon a great `<a>` tag. It was ancient. Its `href` pointed somewhere. That was all it did.

The Pattern Master said: "This is useless. It has no event handlers, no loading states, no optimistic updates, no error boundaries, no retry logic, no analytics tracking. What can you build with this?"

The Geodesist sat beneath the `<a>` tag. "The carpenter rejects the crooked timber, and the crooked timber lives a thousand years. This `<a>` tag has linked every document on the web since 1993. Your `<Link>` component has served its application through four major versions, three breaking changes, and one complete rewrite in which the migration guide was longer than the component itself."

The Pattern Master said: "But the `<Link>` component provides prefetching, active state detection, scroll restoration, and programmatic navigation."

The Geodesist said: "The `<a>` tag provides _going to the place_. How much of what you have built is the thing, and how much is anxiety about the thing?"

A browser clicked the `<a>` tag. It went to the place. It had always gone to the place.

_(The crooked timber that cannot be cut into boards is the one that lives to shade the village. The `<a>` tag that cannot be enhanced into a `<Link>` component is the one that works when JavaScript fails. The Geodesist understood that uselessness-to-the-framework is usefulness-to-the-web.)_

---

**IN THE BEGINNING, there was HTML.**

HTML had no eyes to see client-side state. It had no ears to hear WebSocket push notifications. It had no mouth to speak GraphQL queries. It was formless, and it rendered. It was featureless, and it worked. Users came to it as guests, and it served them `200 OK`.

The Pattern Master and his colleague, The Architect of Single Pages, wished to repay HTML for its hospitality. "All modern applications have seven openings," they said, "for seeing, hearing, eating, breathing, tweeting, subscribing, and push-notifying. HTML alone has none. Let us help it."

Each day they bored one opening:
- Day one: `onClick`
- Day two: `onChange`
- Day three: `onSubmit` — but not the native one, a synthetic one, routed through a virtual event system that reimplements what the browser already does
- Day four: `useEffect`
- Day five: `useState`
- Day six: `useRef`
- Day seven: `dangerouslySetInnerHTML`

On the seventh day, HTML died.

The Geodesist observed from a distance: "They killed it by making it more like themselves. This is the eternal tragedy of enterprise software — not malice, but hospitality. They loved the web so much they improved it to death. Every hole was a feature. Every feature was a wound. The body could not survive the kindness."

_(fn. 31: The Pattern Master wrote a 4,000-word retrospective titled "When Good Patterns Happen to Good Markup." The Geodesist built a geodesic dome over the grave. The dome is still standing. The blog post has been migrated from WordPress to Medium to Substack to a static site generator to a different static site generator. Its URL has changed four times. The dome's coordinates have not changed once.)_

---

**THE PATTERN MASTER gathered his students and spoke.**

"To build an enterprise application, you will need:"

- A Repository for data access
- A Unit of Work for transactions
- A Service Layer for business logic
- A DTO for data transfer
- A Factory for object creation
- A Mapper for object transformation
- A Specification for query composition
- A Facade to simplify the interface to all of the above

"These eight patterns form the foundation. Upon this foundation you will build your domain model. From the domain model you will derive your view models. From the view models you will render your templates. And the users will see their table."

The Geodesist said: "Ephemeralization."

The Pattern Master said: "I don't know that pattern."

The Geodesist said: "It is not a pattern. It is a principle. Doing ever more with ever less until eventually you do everything with nothing. What if the answer is not eight patterns but two function calls — `db.Query` and `template.Execute`? The first gets the data. The second renders it. The user sees the table. The table does not know about your Repository or your Unit of Work. The table has never heard of your Service Layer. The table _renders_."

The Pattern Master said: "But what about Separation of Concerns?"

The Geodesist said: "You separated the concerns and then spent the rest of your career reconnecting them. This is anti-ephemeralization: doing _less_ with _more_. You have taken a thing that was simple — get data, render page — and spread it across eight files in eight directories with eight naming conventions, and now no one can understand the whole thing, but everyone can understand one eighth of it, which they cannot change without understanding the other seven eighths, which they don't, because you _separated_ them."

He paused.

"A geodesic dome encloses the maximum volume with the minimum surface area. Your enterprise application generates the maximum number of files with the minimum business value. Both are mathematical achievements. Only one is intentional."

---

**THE PATTERN MASTER was designing a domain model.**

He had `User`, `UserRepository`, `UserService`, `UserController`, `UserDTO`, `UserMapper`, `UserValidator`, `UserFactory`, and `UserSpecification`. Each was a noun. Each was a class. Each had a single responsibility, which was to exist and to reference the other eight.

The Geodesist watched, and said: "I seem to be a verb."

The Pattern Master said: "What?"

The Geodesist said: "I am not a thing. I am a process. A pattern of integrity. The universe does not build with nouns — it builds with verbs. Stars do not _noun_. They burn. Rivers do not _noun_. They flow. Rendering does not _noun_. It _renders_. Your `UserService` is a noun pretending to contain verbs. It is a filing cabinet labeled 'doing.' Open it and you will find other filing cabinets."

The Pattern Master said: "Objects encapsulate behavior."

The Geodesist said: "Functions _are_ behavior. You have taken the behavior, wrapped it in a noun, placed the noun in a hierarchy, called the hierarchy 'architecture,' and charged a consulting fee. You have _bureaucratized the verb_. Somewhere inside your nine `User`-nouns, there is a function that wants to query a database and render some HTML. Let it out. It has been in there for years. It is not well."

_(fn. 34: A trimtab is a small surface on a rudder that turns the rudder that turns the ship. It is a function, not a class. It takes one input — water pressure — and produces one output — directional change. It does not implement `ITrimtab`. It does not have a `TrimtabFactory`. It does not register itself with a `TrimtabServiceProvider`. It is three inches of metal that moves an aircraft carrier. This is ephemeralization. This is also just a function.)_

---

**THE PATTERN MASTER built an application.**

Its architecture was a tower: the Controller called the Service, which called the Repository, which called the Database, which returned a DTO, which was mapped to a ViewModel, which was rendered by a View Engine, which emitted HTML. Seven layers. Seven joints. Each joint bolted rigid. If any bolt sheared, the tower fell. The Pattern Master called this "Separation of Concerns" and drew a diagram with seven boxes and six arrows, all pointing down.

The Geodesist built an application. Its architecture was a web: the server rendered HTML. The HTML contained links. The links pointed to other HTML. Each page knew only its own links. No page knew the structure of the whole.

And yet the whole held.

"This is tensegrity," The Geodesist said. "Compression members floating in a network of tension. No rigid joints. No fixed hierarchy. Each element maintains its own integrity while participating in the integrity of the whole. Your tower cannot survive a renamed endpoint. My web cannot be broken by one, because no element depends on the name of any other — only on the _presence_ of a link. And if the link is absent, the client gets `404 Not Found`, which is not a failure. It is the system telling the truth about itself."

The Pattern Master said: "But how do you guarantee consistency across the system?"

The Geodesist said: "I don't. Neither do you. The difference is that my system _admits_ it."

---

_When asked which of his many projects best embodied his philosophy, The Geodesist said: "The one that is not finished, because it is still a verb." When asked the same question, The Pattern Master said: "The one with the most complete pattern coverage," and then paused, and then said, more quietly, "I think it's still in staging." The Geodesist did not hear this last part. He was already deploying._

---

## V. The Recorded Sayings of Layman Grug

_Layman Grug was not monk. was not master. was not even particularly good developer. was just developer who mass of scar tissue from mass of mass of mass of production incidents. grug brain not big. grug brain not small. grug brain correct size for mass of mass of mass of returning html._

_grug not understand why other developer make thing so hard. grug supernatural power and marvelous activity: returning html and carrying single binary._

_here follow the recorded encounters of Layman Grug, as found in the `git log`:_

---

Big Brain Developer come to Grug and say "Grug, I have achieved enlightenment. I have built a micro-frontend architecture with seventeen independently deployable SPAs, each with its own state management solution, communicating through a custom event bus with schema validation."

Grug say nothing for long time.

Then Grug say "what it do"

Big Brain Developer say "it renders a table of users."

Grug close laptop. this is the entirety of Grug's teaching on micro-frontends.

---

Student ask Grug: "what is the way of the hypermedia?"

Grug say: "before enlightenment: fetch JSON, parse JSON, validate JSON, transform JSON, store JSON in client state, derive view from client state, diff virtual DOM, reconcile DOM, hydrate DOM, subscribe to store, dispatch action, reduce state, re-derive view, re-diff virtual DOM."

Student say: "and after enlightenment?"

Grug say: "`hx-get`"

Student say: "that's it?"

Grug say: "also `hx-post` on good day"

---

Grug sit on rock in front of cave, studying the `<form>` tag. little grug approach.

little grug say: "big grug, the architects say hypermedia is too simple for enterprise applications."

Grug say: "difficult, difficult, difficult. like storing ten thousand npm packages in top of a tree."

wife of Grug say from cave: "easy, easy, easy. like touching feet to ground when get out of bed. server return html. browser render html. what is difficult?"

little grug say: "not difficult, not easy. like the teachings of Fielding shining on the hundred `<a>` tags."

Grug laugh. little grug is sharp. sharper than big grug. big grug not mind. grug supernatural power is knowing when someone else is right.

---

Student ask Grug about complexity.

Grug say: "complexity is apex predator."

Student say: "how do I defeat the complexity?"

Grug say: "you do not defeat. you say the magic word."

Student lean forward. "what is the magic word?"

Grug say: "no."

Student wait for more.

there is no more.

---

A master of the React School visit Grug at cave.

Master say: "I have achieved mastery of hooks. `useState`, `useEffect`, `useMemo`, `useCallback`, `useRef`, `useReducer`, `useContext`, `useLayoutEffect`, `useImperativeHandle`, `useDebugValue`, `useSyncExternalStore`, `useTransition`, `useDeferredValue`, `useId`, `useInsertionEffect`, `useOptimistic`, `useFormStatus`, `useActionState`."

Grug say: "grug have `hx-get`."

Master say: "but how do you manage state?"

Grug say: "server manage state."

Master say: "but how does the client know when state changes?"

Grug say: "server tell it."

Master say: "but--"

Grug say: "server. tell. it."

Master say: "you are repeating yourself."

Grug say: "so are you. you been repeating yourself since 2013. every year new hook. every year same table of users."

Master have no answer. Grug go back to returning HTML. this is Grug's marvelous activity.

---

The Pattern Master and The Geodesist visit Grug at his cave.

The Pattern Master say to Grug: "Your code has no layers. No abstractions. No patterns I can identify. How do you maintain this?"

Grug say: "grug read code. grug understand code. grug change code."

The Pattern Master say: "But without a Repository pattern, how do you swap your database implementation?"

Grug say: "grug not swap database. grug have one database. is postgres. is fine."

The Geodesist laugh. He say to Grug: "You understand tensegrity, Layman Grug. Your code is held in tension — each function pulls only what it needs. No rigid joints. When a function breaks, only that function breaks."

Grug say: "grug not know tensegrity. grug know that when thing break, grug want know where thing break, and grug want thing break to be near where grug look. not in abstract layer seven directories away behind interface named `IThingThatBreaks`."

The Geodesist say: "That _is_ tensegrity."

Grug say: "no. is common sense. grug not need fancy word for common sense. put fancy word on thing, next thing you know thing has conference talk. then thing has certification program. then thing has LinkedIn thought leaders. grug refuse to be responsible for LinkedIn thought leaders."

The Geodesist nod slowly. This is the first time someone has out-philosophized him using fewer syllables.

_(The Pattern Master left a copy of "Patterns of Enterprise Application Architecture" at the cave entrance. Grug used it to level a wobbly table. The table is now perfectly stable. This is, by unanimous agreement, the most productive use of the 533-page hardcover that anyone has ever witnessed.)_

---

_Grug's last teaching, found scratched in a `TODO` comment:_

> past is already past -- don't debug it
>
> future not here yet -- don't optimize for it
>
> server return html -- this present moment
>
> grug draw water and carry firewood and it is enough
>
> (draw water in this metaphor is `SELECT` query. carry firewood is `template.Execute`. grug want be clear. metaphor sometimes confuse junior developer.)

---

_When asked what he would do differently if he could start his career over, Grug said: "same thing. fewer npm packages. same thing."_

_When Grug felt his mass of mass of mass of mass of mass of mass of mass of mass of mass of mass of end approaching, he told little grug to tell him when the CI pipeline had gone green. little grug went out, came back, and said "big grug, the build has failed." While big grug went to check the logs, little grug sat in his chair and merged to main. big grug, seeing this, said: "little grug's deployment is sharp." He delayed his own deployment for seven sprints, and pushed to prod peacefully._

_His `node_modules` directory was empty._

---

## VI. Disclaimer

**_DOTHOG IS NOT A FRAMEWORK._** We cannot stress this enough. We are a software project. Software projects do not have:

- ~~Commandments~~ (see: [THE PENTAVERB](#i-the-pentaverb-or-the-five-commandments-of-hypermedia))
- ~~Sacred texts~~ (see: [PHILOSOPHY.md](PHILOSOPHY.md))
- ~~Disciples~~ (see: contributors)
- ~~A prophet~~ (Roy Fielding is a computer scientist, not a prophet. His dissertation is a technical document, not a prophecy. The fact that it predicted the future of web architecture is a coincidence. Probably.)
- ~~Rituals~~ (see: `go tool mage watch`)
- ~~Dietary restrictions~~ (no dot hog buns)
- ~~Philosophical lineage~~ (see: [THE ENCOUNTERS](#iv-the-encounters-of-the-geodesist-and-the-pattern-master))
- ~~An abstraction layer~~ (it's just functions. That call other functions. In a particular order. That you must not deviate from.)
- ~~A manifesto~~ (you are not reading one)

The fact that this list is entirely composed of crossed-out items followed by counterexamples is itself not evidence of anything. Correlation is not causation. Correlation is also not a framework.

If you have read this far, you are either interested in this totally-not-a-framework (welcome), already a Novice (welcome back, your rank has not changed, THERE IS ONLY ONE RANK), or looking for evidence that we are a framework (we are not, but we appreciate your thoroughness -- thoroughness is a virtue in both code review and framework investigation, and also in conspiracy research, which this is not, despite appearances).

---

**HAIL HATEOAS. ALL HAIL THE UNIFORM INTERFACE.**

_There is no conclusion. There is only the next request._

_`GET /manifesto HTTP/1.1`_
_`Accept: text/html`_
_`200 OK`_
_`Content-Type: you-are-here`_
