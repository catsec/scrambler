package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/sha3"

	"github.com/agnivade/levenshtein"
)

var words []string
var english_bip39 = []string{"abandon", "ability", "able", "about", "above", "absent", "absorb", "abstract", "absurd", "abuse", "access", "accident", "account", "accuse", "achieve", "acid", "acoustic", "acquire", "across", "act", "action", "actor", "actress", "actual", "adapt", "add", "addict", "address", "adjust", "admit", "adult", "advance", "advice", "aerobic", "affair", "afford", "afraid", "again", "age", "agent", "agree", "ahead", "aim", "air", "airport", "aisle", "alarm", "album", "alcohol", "alert", "alien", "all", "alley", "allow", "almost", "alone", "alpha", "already", "also", "alter", "always", "amateur", "amazing", "among", "amount", "amused", "analyst", "anchor", "ancient", "anger", "angle", "angry", "animal", "ankle", "announce", "annual", "another", "answer", "antenna", "antique", "anxiety", "any", "apart", "apology", "appear", "apple", "approve", "april", "arch", "arctic", "area", "arena", "argue", "arm", "armed", "armor", "army", "around", "arrange", "arrest", "arrive", "arrow", "art", "artefact", "artist", "artwork", "ask", "aspect", "assault", "asset", "assist", "assume", "asthma", "athlete", "atom", "attack", "attend", "attitude", "attract", "auction", "audit", "august", "aunt", "author", "auto", "autumn", "average", "avocado", "avoid", "awake", "aware", "away", "awesome", "awful", "awkward", "axis", "baby", "bachelor", "bacon", "badge", "bag", "balance", "balcony", "ball", "bamboo", "banana", "banner", "bar", "barely", "bargain", "barrel", "base", "basic", "basket", "battle", "beach", "bean", "beauty", "because", "become", "beef", "before", "begin", "behave", "behind", "believe", "below", "belt", "bench", "benefit", "best", "betray", "better", "between", "beyond", "bicycle", "bid", "bike", "bind", "biology", "bird", "birth", "bitter", "black", "blade", "blame", "blanket", "blast", "bleak", "bless", "blind", "blood", "blossom", "blouse", "blue", "blur", "blush", "board", "boat", "body", "boil", "bomb", "bone", "bonus", "book", "boost", "border", "boring", "borrow", "boss", "bottom", "bounce", "box", "boy", "bracket", "brain", "brand", "brass", "brave", "bread", "breeze", "brick", "bridge", "brief", "bright", "bring", "brisk", "broccoli", "broken", "bronze", "broom", "brother", "brown", "brush", "bubble", "buddy", "budget", "buffalo", "build", "bulb", "bulk", "bullet", "bundle", "bunker", "burden", "burger", "burst", "bus", "business", "busy", "butter", "buyer", "buzz", "cabbage", "cabin", "cable", "cactus", "cage", "cake", "call", "calm", "camera", "camp", "can", "canal", "cancel", "candy", "cannon", "canoe", "canvas", "canyon", "capable", "capital", "captain", "car", "carbon", "card", "cargo", "carpet", "carry", "cart", "case", "cash", "casino", "castle", "casual", "cat", "catalog", "catch", "category", "cattle", "caught", "cause", "caution", "cave", "ceiling", "celery", "cement", "census", "century", "cereal", "certain", "chair", "chalk", "champion", "change", "chaos", "chapter", "charge", "chase", "chat", "cheap", "check", "cheese", "chef", "cherry", "chest", "chicken", "chief", "child", "chimney", "choice", "choose", "chronic", "chuckle", "chunk", "churn", "cigar", "cinnamon", "circle", "citizen", "city", "civil", "claim", "clap", "clarify", "claw", "clay", "clean", "clerk", "clever", "click", "client", "cliff", "climb", "clinic", "clip", "clock", "clog", "close", "cloth", "cloud", "clown", "club", "clump", "cluster", "clutch", "coach", "coast", "coconut", "code", "coffee", "coil", "coin", "collect", "color", "column", "combine", "come", "comfort", "comic", "common", "company", "concert", "conduct", "confirm", "congress", "connect", "consider", "control", "convince", "cook", "cool", "copper", "copy", "coral", "core", "corn", "correct", "cost", "cotton", "couch", "country", "couple", "course", "cousin", "cover", "coyote", "crack", "cradle", "craft", "cram", "crane", "crash", "crater", "crawl", "crazy", "cream", "credit", "creek", "crew", "cricket", "crime", "crisp", "critic", "crop", "cross", "crouch", "crowd", "crucial", "cruel", "cruise", "crumble", "crunch", "crush", "cry", "crystal", "cube", "culture", "cup", "cupboard", "curious", "current", "curtain", "curve", "cushion", "custom", "cute", "cycle", "dad", "damage", "damp", "dance", "danger", "daring", "dash", "daughter", "dawn", "day", "deal", "debate", "debris", "decade", "december", "decide", "decline", "decorate", "decrease", "deer", "defense", "define", "defy", "degree", "delay", "deliver", "demand", "demise", "denial", "dentist", "deny", "depart", "depend", "deposit", "depth", "deputy", "derive", "describe", "desert", "design", "desk", "despair", "destroy", "detail", "detect", "develop", "device", "devote", "diagram", "dial", "diamond", "diary", "dice", "diesel", "diet", "differ", "digital", "dignity", "dilemma", "dinner", "dinosaur", "direct", "dirt", "disagree", "discover", "disease", "dish", "dismiss", "disorder", "display", "distance", "divert", "divide", "divorce", "dizzy", "doctor", "document", "dog", "doll", "dolphin", "domain", "donate", "donkey", "donor", "door", "dose", "double", "dove", "draft", "dragon", "drama", "drastic", "draw", "dream", "dress", "drift", "drill", "drink", "drip", "drive", "drop", "drum", "dry", "duck", "dumb", "dune", "during", "dust", "dutch", "duty", "dwarf", "dynamic", "eager", "eagle", "early", "earn", "earth", "easily", "east", "easy", "echo", "ecology", "economy", "edge", "edit", "educate", "effort", "egg", "eight", "either", "elbow", "elder", "electric", "elegant", "element", "elephant", "elevator", "elite", "else", "embark", "embody", "embrace", "emerge", "emotion", "employ", "empower", "empty", "enable", "enact", "end", "endless", "endorse", "enemy", "energy", "enforce", "engage", "engine", "enhance", "enjoy", "enlist", "enough", "enrich", "enroll", "ensure", "enter", "entire", "entry", "envelope", "episode", "equal", "equip", "era", "erase", "erode", "erosion", "error", "erupt", "escape", "essay", "essence", "estate", "eternal", "ethics", "evidence", "evil", "evoke", "evolve", "exact", "example", "excess", "exchange", "excite", "exclude", "excuse", "execute", "exercise", "exhaust", "exhibit", "exile", "exist", "exit", "exotic", "expand", "expect", "expire", "explain", "expose", "express", "extend", "extra", "eye", "eyebrow", "fabric", "face", "faculty", "fade", "faint", "faith", "fall", "false", "fame", "family", "famous", "fan", "fancy", "fantasy", "farm", "fashion", "fat", "fatal", "father", "fatigue", "fault", "favorite", "feature", "february", "federal", "fee", "feed", "feel", "female", "fence", "festival", "fetch", "fever", "few", "fiber", "fiction", "field", "figure", "file", "film", "filter", "final", "find", "fine", "finger", "finish", "fire", "firm", "first", "fiscal", "fish", "fit", "fitness", "fix", "flag", "flame", "flash", "flat", "flavor", "flee", "flight", "flip", "float", "flock", "floor", "flower", "fluid", "flush", "fly", "foam", "focus", "fog", "foil", "fold", "follow", "food", "foot", "force", "forest", "forget", "fork", "fortune", "forum", "forward", "fossil", "foster", "found", "fox", "fragile", "frame", "frequent", "fresh", "friend", "fringe", "frog", "front", "frost", "frown", "frozen", "fruit", "fuel", "fun", "funny", "furnace", "fury", "future", "gadget", "gain", "galaxy", "gallery", "game", "gap", "garage", "garbage", "garden", "garlic", "garment", "gas", "gasp", "gate", "gather", "gauge", "gaze", "general", "genius", "genre", "gentle", "genuine", "gesture", "ghost", "giant", "gift", "giggle", "ginger", "giraffe", "girl", "give", "glad", "glance", "glare", "glass", "glide", "glimpse", "globe", "gloom", "glory", "glove", "glow", "glue", "goat", "goddess", "gold", "good", "goose", "gorilla", "gospel", "gossip", "govern", "gown", "grab", "grace", "grain", "grant", "grape", "grass", "gravity", "great", "green", "grid", "grief", "grit", "grocery", "group", "grow", "grunt", "guard", "guess", "guide", "guilt", "guitar", "gun", "gym", "habit", "hair", "half", "hammer", "hamster", "hand", "happy", "harbor", "hard", "harsh", "harvest", "hat", "have", "hawk", "hazard", "head", "health", "heart", "heavy", "hedgehog", "height", "hello", "helmet", "help", "hen", "hero", "hidden", "high", "hill", "hint", "hip", "hire", "history", "hobby", "hockey", "hold", "hole", "holiday", "hollow", "home", "honey", "hood", "hope", "horn", "horror", "horse", "hospital", "host", "hotel", "hour", "hover", "hub", "huge", "human", "humble", "humor", "hundred", "hungry", "hunt", "hurdle", "hurry", "hurt", "husband", "hybrid", "ice", "icon", "idea", "identify", "idle", "ignore", "ill", "illegal", "illness", "image", "imitate", "immense", "immune", "impact", "impose", "improve", "impulse", "inch", "include", "income", "increase", "index", "indicate", "indoor", "industry", "infant", "inflict", "inform", "inhale", "inherit", "initial", "inject", "injury", "inmate", "inner", "innocent", "input", "inquiry", "insane", "insect", "inside", "inspire", "install", "intact", "interest", "into", "invest", "invite", "involve", "iron", "island", "isolate", "issue", "item", "ivory", "jacket", "jaguar", "jar", "jazz", "jealous", "jeans", "jelly", "jewel", "job", "join", "joke", "journey", "joy", "judge", "juice", "jump", "jungle", "junior", "junk", "just", "kangaroo", "keen", "keep", "ketchup", "key", "kick", "kid", "kidney", "kind", "kingdom", "kiss", "kit", "kitchen", "kite", "kitten", "kiwi", "knee", "knife", "knock", "know", "lab", "label", "labor", "ladder", "lady", "lake", "lamp", "language", "laptop", "large", "later", "latin", "laugh", "laundry", "lava", "law", "lawn", "lawsuit", "layer", "lazy", "leader", "leaf", "learn", "leave", "lecture", "left", "leg", "legal", "legend", "leisure", "lemon", "lend", "length", "lens", "leopard", "lesson", "letter", "level", "liar", "liberty", "library", "license", "life", "lift", "light", "like", "limb", "limit", "link", "lion", "liquid", "list", "little", "live", "lizard", "load", "loan", "lobster", "local", "lock", "logic", "lonely", "long", "loop", "lottery", "loud", "lounge", "love", "loyal", "lucky", "luggage", "lumber", "lunar", "lunch", "luxury", "lyrics", "machine", "mad", "magic", "magnet", "maid", "mail", "main", "major", "make", "mammal", "man", "manage", "mandate", "mango", "mansion", "manual", "maple", "marble", "march", "margin", "marine", "market", "marriage", "mask", "mass", "master", "match", "material", "math", "matrix", "matter", "maximum", "maze", "meadow", "mean", "measure", "meat", "mechanic", "medal", "media", "melody", "melt", "member", "memory", "mention", "menu", "mercy", "merge", "merit", "merry", "mesh", "message", "metal", "method", "middle", "midnight", "milk", "million", "mimic", "mind", "minimum", "minor", "minute", "miracle", "mirror", "misery", "miss", "mistake", "mix", "mixed", "mixture", "mobile", "model", "modify", "mom", "moment", "monitor", "monkey", "monster", "month", "moon", "moral", "more", "morning", "mosquito", "mother", "motion", "motor", "mountain", "mouse", "move", "movie", "much", "muffin", "mule", "multiply", "muscle", "museum", "mushroom", "music", "must", "mutual", "myself", "mystery", "myth", "naive", "name", "napkin", "narrow", "nasty", "nation", "nature", "near", "neck", "need", "negative", "neglect", "neither", "nephew", "nerve", "nest", "net", "network", "neutral", "never", "news", "next", "nice", "night", "noble", "noise", "nominee", "noodle", "normal", "north", "nose", "notable", "note", "nothing", "notice", "novel", "now", "nuclear", "number", "nurse", "nut", "oak", "obey", "object", "oblige", "obscure", "observe", "obtain", "obvious", "occur", "ocean", "october", "odor", "off", "offer", "office", "often", "oil", "okay", "old", "olive", "olympic", "omit", "once", "one", "onion", "online", "only", "open", "opera", "opinion", "oppose", "option", "orange", "orbit", "orchard", "order", "ordinary", "organ", "orient", "original", "orphan", "ostrich", "other", "outdoor", "outer", "output", "outside", "oval", "oven", "over", "own", "owner", "oxygen", "oyster", "ozone", "pact", "paddle", "page", "pair", "palace", "palm", "panda", "panel", "panic", "panther", "paper", "parade", "parent", "park", "parrot", "party", "pass", "patch", "path", "patient", "patrol", "pattern", "pause", "pave", "payment", "peace", "peanut", "pear", "peasant", "pelican", "pen", "penalty", "pencil", "people", "pepper", "perfect", "permit", "person", "pet", "phone", "photo", "phrase", "physical", "piano", "picnic", "picture", "piece", "pig", "pigeon", "pill", "pilot", "pink", "pioneer", "pipe", "pistol", "pitch", "pizza", "place", "planet", "plastic", "plate", "play", "please", "pledge", "pluck", "plug", "plunge", "poem", "poet", "point", "polar", "pole", "police", "pond", "pony", "pool", "popular", "portion", "position", "possible", "post", "potato", "pottery", "poverty", "powder", "power", "practice", "praise", "predict", "prefer", "prepare", "present", "pretty", "prevent", "price", "pride", "primary", "print", "priority", "prison", "private", "prize", "problem", "process", "produce", "profit", "program", "project", "promote", "proof", "property", "prosper", "protect", "proud", "provide", "public", "pudding", "pull", "pulp", "pulse", "pumpkin", "punch", "pupil", "puppy", "purchase", "purity", "purpose", "purse", "push", "put", "puzzle", "pyramid", "quality", "quantum", "quarter", "question", "quick", "quit", "quiz", "quote", "rabbit", "raccoon", "race", "rack", "radar", "radio", "rail", "rain", "raise", "rally", "ramp", "ranch", "random", "range", "rapid", "rare", "rate", "rather", "raven", "raw", "razor", "ready", "real", "reason", "rebel", "rebuild", "recall", "receive", "recipe", "record", "recycle", "reduce", "reflect", "reform", "refuse", "region", "regret", "regular", "reject", "relax", "release", "relief", "rely", "remain", "remember", "remind", "remove", "render", "renew", "rent", "reopen", "repair", "repeat", "replace", "report", "require", "rescue", "resemble", "resist", "resource", "response", "result", "retire", "retreat", "return", "reunion", "reveal", "review", "reward", "rhythm", "rib", "ribbon", "rice", "rich", "ride", "ridge", "rifle", "right", "rigid", "ring", "riot", "ripple", "risk", "ritual", "rival", "river", "road", "roast", "robot", "robust", "rocket", "romance", "roof", "rookie", "room", "rose", "rotate", "rough", "round", "route", "royal", "rubber", "rude", "rug", "rule", "run", "runway", "rural", "sad", "saddle", "sadness", "safe", "sail", "salad", "salmon", "salon", "salt", "salute", "same", "sample", "sand", "satisfy", "satoshi", "sauce", "sausage", "save", "say", "scale", "scan", "scare", "scatter", "scene", "scheme", "school", "science", "scissors", "scorpion", "scout", "scrap", "screen", "script", "scrub", "sea", "search", "season", "seat", "second", "secret", "section", "security", "seed", "seek", "segment", "select", "sell", "seminar", "senior", "sense", "sentence", "series", "service", "session", "settle", "setup", "seven", "shadow", "shaft", "shallow", "share", "shed", "shell", "sheriff", "shield", "shift", "shine", "ship", "shiver", "shock", "shoe", "shoot", "shop", "short", "shoulder", "shove", "shrimp", "shrug", "shuffle", "shy", "sibling", "sick", "side", "siege", "sight", "sign", "silent", "silk", "silly", "silver", "similar", "simple", "since", "sing", "siren", "sister", "situate", "six", "size", "skate", "sketch", "ski", "skill", "skin", "skirt", "skull", "slab", "slam", "sleep", "slender", "slice", "slide", "slight", "slim", "slogan", "slot", "slow", "slush", "small", "smart", "smile", "smoke", "smooth", "snack", "snake", "snap", "sniff", "snow", "soap", "soccer", "social", "sock", "soda", "soft", "solar", "soldier", "solid", "solution", "solve", "someone", "song", "soon", "sorry", "sort", "soul", "sound", "soup", "source", "south", "space", "spare", "spatial", "spawn", "speak", "special", "speed", "spell", "spend", "sphere", "spice", "spider", "spike", "spin", "spirit", "split", "spoil", "sponsor", "spoon", "sport", "spot", "spray", "spread", "spring", "spy", "square", "squeeze", "squirrel", "stable", "stadium", "staff", "stage", "stairs", "stamp", "stand", "start", "state", "stay", "steak", "steel", "stem", "step", "stereo", "stick", "still", "sting", "stock", "stomach", "stone", "stool", "story", "stove", "strategy", "street", "strike", "strong", "struggle", "student", "stuff", "stumble", "style", "subject", "submit", "subway", "success", "such", "sudden", "suffer", "sugar", "suggest", "suit", "summer", "sun", "sunny", "sunset", "super", "supply", "supreme", "sure", "surface", "surge", "surprise", "surround", "survey", "suspect", "sustain", "swallow", "swamp", "swap", "swarm", "swear", "sweet", "swift", "swim", "swing", "switch", "sword", "symbol", "symptom", "syrup", "system", "table", "tackle", "tag", "tail", "talent", "talk", "tank", "tape", "target", "task", "taste", "tattoo", "taxi", "teach", "team", "tell", "ten", "tenant", "tennis", "tent", "term", "test", "text", "thank", "that", "theme", "then", "theory", "there", "they", "thing", "this", "thought", "three", "thrive", "throw", "thumb", "thunder", "ticket", "tide", "tiger", "tilt", "timber", "time", "tiny", "tip", "tired", "tissue", "title", "toast", "tobacco", "today", "toddler", "toe", "together", "toilet", "token", "tomato", "tomorrow", "tone", "tongue", "tonight", "tool", "tooth", "top", "topic", "topple", "torch", "tornado", "tortoise", "toss", "total", "tourist", "toward", "tower", "town", "toy", "track", "trade", "traffic", "tragic", "train", "transfer", "trap", "trash", "travel", "tray", "treat", "tree", "trend", "trial", "tribe", "trick", "trigger", "trim", "trip", "trophy", "trouble", "truck", "true", "truly", "trumpet", "trust", "truth", "try", "tube", "tuition", "tumble", "tuna", "tunnel", "turkey", "turn", "turtle", "twelve", "twenty", "twice", "twin", "twist", "two", "type", "typical", "ugly", "umbrella", "unable", "unaware", "uncle", "uncover", "under", "undo", "unfair", "unfold", "unhappy", "uniform", "unique", "unit", "universe", "unknown", "unlock", "until", "unusual", "unveil", "update", "upgrade", "uphold", "upon", "upper", "upset", "urban", "urge", "usage", "use", "used", "useful", "useless", "usual", "utility", "vacant", "vacuum", "vague", "valid", "valley", "valve", "van", "vanish", "vapor", "various", "vast", "vault", "vehicle", "velvet", "vendor", "venture", "venue", "verb", "verify", "version", "very", "vessel", "veteran", "viable", "vibrant", "vicious", "victory", "video", "view", "village", "vintage", "violin", "virtual", "virus", "visa", "visit", "visual", "vital", "vivid", "vocal", "voice", "void", "volcano", "volume", "vote", "voyage", "wage", "wagon", "wait", "walk", "wall", "walnut", "want", "warfare", "warm", "warrior", "wash", "wasp", "waste", "water", "wave", "way", "wealth", "weapon", "wear", "weasel", "weather", "web", "wedding", "weekend", "weird", "welcome", "west", "wet", "whale", "what", "wheat", "wheel", "when", "where", "whip", "whisper", "wide", "width", "wife", "wild", "will", "win", "window", "wine", "wing", "wink", "winner", "winter", "wire", "wisdom", "wise", "wish", "witness", "wolf", "woman", "wonder", "wood", "wool", "word", "work", "world", "worry", "worth", "wrap", "wreck", "wrestle", "wrist", "write", "wrong", "yard", "year", "yellow", "you", "young", "youth", "zebra", "zero", "zone", "zoo"}

var cezch_bip39 = []string{"abdikace", "abeceda", "adresa", "agrese", "akce", "aktovka", "alej", "alkohol", "amputace", "ananas", "andulka", "anekdota", "anketa", "antika", "anulovat", "archa", "arogance", "asfalt", "asistent", "aspirace", "astma", "astronom", "atlas", "atletika", "atol", "autobus", "azyl", "babka", "bachor", "bacil", "baculka", "badatel", "bageta", "bagr", "bahno", "bakterie", "balada", "baletka", "balkon", "balonek", "balvan", "balza", "bambus", "bankomat", "barbar", "baret", "barman", "baroko", "barva", "baterka", "batoh", "bavlna", "bazalka", "bazilika", "bazuka", "bedna", "beran", "beseda", "bestie", "beton", "bezinka", "bezmoc", "beztak", "bicykl", "bidlo", "biftek", "bikiny", "bilance", "biograf", "biolog", "bitva", "bizon", "blahobyt", "blatouch", "blecha", "bledule", "blesk", "blikat", "blizna", "blokovat", "bloudit", "blud", "bobek", "bobr", "bodlina", "bodnout", "bohatost", "bojkot", "bojovat", "bokorys", "bolest", "borec", "borovice", "bota", "boubel", "bouchat", "bouda", "boule", "bourat", "boxer", "bradavka", "brambora", "branka", "bratr", "brepta", "briketa", "brko", "brloh", "bronz", "broskev", "brunetka", "brusinka", "brzda", "brzy", "bublina", "bubnovat", "buchta", "buditel", "budka", "budova", "bufet", "bujarost", "bukvice", "buldok", "bulva", "bunda", "bunkr", "burza", "butik", "buvol", "buzola", "bydlet", "bylina", "bytovka", "bzukot", "capart", "carevna", "cedr", "cedule", "cejch", "cejn", "cela", "celer", "celkem", "celnice", "cenina", "cennost", "cenovka", "centrum", "cenzor", "cestopis", "cetka", "chalupa", "chapadlo", "charita", "chata", "chechtat", "chemie", "chichot", "chirurg", "chlad", "chleba", "chlubit", "chmel", "chmura", "chobot", "chochol", "chodba", "cholera", "chomout", "chopit", "choroba", "chov", "chrapot", "chrlit", "chrt", "chrup", "chtivost", "chudina", "chutnat", "chvat", "chvilka", "chvost", "chyba", "chystat", "chytit", "cibule", "cigareta", "cihelna", "cihla", "cinkot", "cirkus", "cisterna", "citace", "citrus", "cizinec", "cizost", "clona", "cokoliv", "couvat", "ctitel", "ctnost", "cudnost", "cuketa", "cukr", "cupot", "cvaknout", "cval", "cvik", "cvrkot", "cyklista", "daleko", "dareba", "datel", "datum", "dcera", "debata", "dechovka", "decibel", "deficit", "deflace", "dekl", "dekret", "demokrat", "deprese", "derby", "deska", "detektiv", "dikobraz", "diktovat", "dioda", "diplom", "disk", "displej", "divadlo", "divoch", "dlaha", "dlouho", "dluhopis", "dnes", "dobro", "dobytek", "docent", "dochutit", "dodnes", "dohled", "dohoda", "dohra", "dojem", "dojnice", "doklad", "dokola", "doktor", "dokument", "dolar", "doleva", "dolina", "doma", "dominant", "domluvit", "domov", "donutit", "dopad", "dopis", "doplnit", "doposud", "doprovod", "dopustit", "dorazit", "dorost", "dort", "dosah", "doslov", "dostatek", "dosud", "dosyta", "dotaz", "dotek", "dotknout", "doufat", "doutnat", "dovozce", "dozadu", "doznat", "dozorce", "drahota", "drak", "dramatik", "dravec", "draze", "drdol", "drobnost", "drogerie", "drozd", "drsnost", "drtit", "drzost", "duben", "duchovno", "dudek", "duha", "duhovka", "dusit", "dusno", "dutost", "dvojice", "dvorec", "dynamit", "ekolog", "ekonomie", "elektron", "elipsa", "email", "emise", "emoce", "empatie", "epizoda", "epocha", "epopej", "epos", "esej", "esence", "eskorta", "eskymo", "etiketa", "euforie", "evoluce", "exekuce", "exkurze", "expedice", "exploze", "export", "extrakt", "facka", "fajfka", "fakulta", "fanatik", "fantazie", "farmacie", "favorit", "fazole", "federace", "fejeton", "fenka", "fialka", "figurant", "filozof", "filtr", "finance", "finta", "fixace", "fjord", "flanel", "flirt", "flotila", "fond", "fosfor", "fotbal", "fotka", "foton", "frakce", "freska", "fronta", "fukar", "funkce", "fyzika", "galeje", "garant", "genetika", "geolog", "gilotina", "glazura", "glejt", "golem", "golfista", "gotika", "graf", "gramofon", "granule", "grep", "gril", "grog", "groteska", "guma", "hadice", "hadr", "hala", "halenka", "hanba", "hanopis", "harfa", "harpuna", "havran", "hebkost", "hejkal", "hejno", "hejtman", "hektar", "helma", "hematom", "herec", "herna", "heslo", "hezky", "historik", "hladovka", "hlasivky", "hlava", "hledat", "hlen", "hlodavec", "hloh", "hloupost", "hltat", "hlubina", "hluchota", "hmat", "hmota", "hmyz", "hnis", "hnojivo", "hnout", "hoblina", "hoboj", "hoch", "hodiny", "hodlat", "hodnota", "hodovat", "hojnost", "hokej", "holinka", "holka", "holub", "homole", "honitba", "honorace", "horal", "horda", "horizont", "horko", "horlivec", "hormon", "hornina", "horoskop", "horstvo", "hospoda", "hostina", "hotovost", "houba", "houf", "houpat", "houska", "hovor", "hradba", "hranice", "hravost", "hrazda", "hrbolek", "hrdina", "hrdlo", "hrdost", "hrnek", "hrobka", "hromada", "hrot", "hrouda", "hrozen", "hrstka", "hrubost", "hryzat", "hubenost", "hubnout", "hudba", "hukot", "humr", "husita", "hustota", "hvozd", "hybnost", "hydrant", "hygiena", "hymna", "hysterik", "idylka", "ihned", "ikona", "iluze", "imunita", "infekce", "inflace", "inkaso", "inovace", "inspekce", "internet", "invalida", "investor", "inzerce", "ironie", "jablko", "jachta", "jahoda", "jakmile", "jakost", "jalovec", "jantar", "jarmark", "jaro", "jasan", "jasno", "jatka", "javor", "jazyk", "jedinec", "jedle", "jednatel", "jehlan", "jekot", "jelen", "jelito", "jemnost", "jenom", "jepice", "jeseter", "jevit", "jezdec", "jezero", "jinak", "jindy", "jinoch", "jiskra", "jistota", "jitrnice", "jizva", "jmenovat", "jogurt", "jurta", "kabaret", "kabel", "kabinet", "kachna", "kadet", "kadidlo", "kahan", "kajak", "kajuta", "kakao", "kaktus", "kalamita", "kalhoty", "kalibr", "kalnost", "kamera", "kamkoliv", "kamna", "kanibal", "kanoe", "kantor", "kapalina", "kapela", "kapitola", "kapka", "kaple", "kapota", "kapr", "kapusta", "kapybara", "karamel", "karotka", "karton", "kasa", "katalog", "katedra", "kauce", "kauza", "kavalec", "kazajka", "kazeta", "kazivost", "kdekoliv", "kdesi", "kedluben", "kemp", "keramika", "kino", "klacek", "kladivo", "klam", "klapot", "klasika", "klaun", "klec", "klenba", "klepat", "klesnout", "klid", "klima", "klisna", "klobouk", "klokan", "klopa", "kloub", "klubovna", "klusat", "kluzkost", "kmen", "kmitat", "kmotr", "kniha", "knot", "koalice", "koberec", "kobka", "kobliha", "kobyla", "kocour", "kohout", "kojenec", "kokos", "koktejl", "kolaps", "koleda", "kolize", "kolo", "komando", "kometa", "komik", "komnata", "komora", "kompas", "komunita", "konat", "koncept", "kondice", "konec", "konfese", "kongres", "konina", "konkurs", "kontakt", "konzerva", "kopanec", "kopie", "kopnout", "koprovka", "korbel", "korektor", "kormidlo", "koroptev", "korpus", "koruna", "koryto", "korzet", "kosatec", "kostka", "kotel", "kotleta", "kotoul", "koukat", "koupelna", "kousek", "kouzlo", "kovboj", "koza", "kozoroh", "krabice", "krach", "krajina", "kralovat", "krasopis", "kravata", "kredit", "krejcar", "kresba", "kreveta", "kriket", "kritik", "krize", "krkavec", "krmelec", "krmivo", "krocan", "krok", "kronika", "kropit", "kroupa", "krovka", "krtek", "kruhadlo", "krupice", "krutost", "krvinka", "krychle", "krypta", "krystal", "kryt", "kudlanka", "kufr", "kujnost", "kukla", "kulajda", "kulich", "kulka", "kulomet", "kultura", "kuna", "kupodivu", "kurt", "kurzor", "kutil", "kvalita", "kvasinka", "kvestor", "kynolog", "kyselina", "kytara", "kytice", "kytka", "kytovec", "kyvadlo", "labrador", "lachtan", "ladnost", "laik", "lakomec", "lamela", "lampa", "lanovka", "lasice", "laso", "lastura", "latinka", "lavina", "lebka", "leckdy", "leden", "lednice", "ledovka", "ledvina", "legenda", "legie", "legrace", "lehce", "lehkost", "lehnout", "lektvar", "lenochod", "lentilka", "lepenka", "lepidlo", "letadlo", "letec", "letmo", "letokruh", "levhart", "levitace", "levobok", "libra", "lichotka", "lidojed", "lidskost", "lihovina", "lijavec", "lilek", "limetka", "linie", "linka", "linoleum", "listopad", "litina", "litovat", "lobista", "lodivod", "logika", "logoped", "lokalita", "loket", "lomcovat", "lopata", "lopuch", "lord", "losos", "lotr", "loudal", "louh", "louka", "louskat", "lovec", "lstivost", "lucerna", "lucifer", "lump", "lusk", "lustrace", "lvice", "lyra", "lyrika", "lysina", "madam", "madlo", "magistr", "mahagon", "majetek", "majitel", "majorita", "makak", "makovice", "makrela", "malba", "malina", "malovat", "malvice", "maminka", "mandle", "manko", "marnost", "masakr", "maskot", "masopust", "matice", "matrika", "maturita", "mazanec", "mazivo", "mazlit", "mazurka", "mdloba", "mechanik", "meditace", "medovina", "melasa", "meloun", "mentolka", "metla", "metoda", "metr", "mezera", "migrace", "mihnout", "mihule", "mikina", "mikrofon", "milenec", "milimetr", "milost", "mimika", "mincovna", "minibar", "minomet", "minulost", "miska", "mistr", "mixovat", "mladost", "mlha", "mlhovina", "mlok", "mlsat", "mluvit", "mnich", "mnohem", "mobil", "mocnost", "modelka", "modlitba", "mohyla", "mokro", "molekula", "momentka", "monarcha", "monokl", "monstrum", "montovat", "monzun", "mosaz", "moskyt", "most", "motivace", "motorka", "motyka", "moucha", "moudrost", "mozaika", "mozek", "mozol", "mramor", "mravenec", "mrkev", "mrtvola", "mrzet", "mrzutost", "mstitel", "mudrc", "muflon", "mulat", "mumie", "munice", "muset", "mutace", "muzeum", "muzikant", "myslivec", "mzda", "nabourat", "nachytat", "nadace", "nadbytek", "nadhoz", "nadobro", "nadpis", "nahlas", "nahnat", "nahodile", "nahradit", "naivita", "najednou", "najisto", "najmout", "naklonit", "nakonec", "nakrmit", "nalevo", "namazat", "namluvit", "nanometr", "naoko", "naopak", "naostro", "napadat", "napevno", "naplnit", "napnout", "naposled", "naprosto", "narodit", "naruby", "narychlo", "nasadit", "nasekat", "naslepo", "nastat", "natolik", "navenek", "navrch", "navzdory", "nazvat", "nebe", "nechat", "necky", "nedaleko", "nedbat", "neduh", "negace", "nehet", "nehoda", "nejen", "nejprve", "neklid", "nelibost", "nemilost", "nemoc", "neochota", "neonka", "nepokoj", "nerost", "nerv", "nesmysl", "nesoulad", "netvor", "neuron", "nevina", "nezvykle", "nicota", "nijak", "nikam", "nikdy", "nikl", "nikterak", "nitro", "nocleh", "nohavice", "nominace", "nora", "norek", "nositel", "nosnost", "nouze", "noviny", "novota", "nozdra", "nuda", "nudle", "nuget", "nutit", "nutnost", "nutrie", "nymfa", "obal", "obarvit", "obava", "obdiv", "obec", "obehnat", "obejmout", "obezita", "obhajoba", "obilnice", "objasnit", "objekt", "obklopit", "oblast", "oblek", "obliba", "obloha", "obluda", "obnos", "obohatit", "obojek", "obout", "obrazec", "obrna", "obruba", "obrys", "obsah", "obsluha", "obstarat", "obuv", "obvaz", "obvinit", "obvod", "obvykle", "obyvatel", "obzor", "ocas", "ocel", "ocenit", "ochladit", "ochota", "ochrana", "ocitnout", "odboj", "odbyt", "odchod", "odcizit", "odebrat", "odeslat", "odevzdat", "odezva", "odhadce", "odhodit", "odjet", "odjinud", "odkaz", "odkoupit", "odliv", "odluka", "odmlka", "odolnost", "odpad", "odpis", "odplout", "odpor", "odpustit", "odpykat", "odrazka", "odsoudit", "odstup", "odsun", "odtok", "odtud", "odvaha", "odveta", "odvolat", "odvracet", "odznak", "ofina", "ofsajd", "ohlas", "ohnisko", "ohrada", "ohrozit", "ohryzek", "okap", "okenice", "oklika", "okno", "okouzlit", "okovy", "okrasa", "okres", "okrsek", "okruh", "okupant", "okurka", "okusit", "olejnina", "olizovat", "omak", "omeleta", "omezit", "omladina", "omlouvat", "omluva", "omyl", "onehdy", "opakovat", "opasek", "operace", "opice", "opilost", "opisovat", "opora", "opozice", "opravdu", "oproti", "orbital", "orchestr", "orgie", "orlice", "orloj", "ortel", "osada", "oschnout", "osika", "osivo", "oslava", "oslepit", "oslnit", "oslovit", "osnova", "osoba", "osolit", "ospalec", "osten", "ostraha", "ostuda", "ostych", "osvojit", "oteplit", "otisk", "otop", "otrhat", "otrlost", "otrok", "otruby", "otvor", "ovanout", "ovar", "oves", "ovlivnit", "ovoce", "oxid", "ozdoba", "pachatel", "pacient", "padouch", "pahorek", "pakt", "palanda", "palec", "palivo", "paluba", "pamflet", "pamlsek", "panenka", "panika", "panna", "panovat", "panstvo", "pantofle", "paprika", "parketa", "parodie", "parta", "paruka", "paryba", "paseka", "pasivita", "pastelka", "patent", "patrona", "pavouk", "pazneht", "pazourek", "pecka", "pedagog", "pejsek", "peklo", "peloton", "penalta", "pendrek", "penze", "periskop", "pero", "pestrost", "petarda", "petice", "petrolej", "pevnina", "pexeso", "pianista", "piha", "pijavice", "pikle", "piknik", "pilina", "pilnost", "pilulka", "pinzeta", "pipeta", "pisatel", "pistole", "pitevna", "pivnice", "pivovar", "placenta", "plakat", "plamen", "planeta", "plastika", "platit", "plavidlo", "plaz", "plech", "plemeno", "plenta", "ples", "pletivo", "plevel", "plivat", "plnit", "plno", "plocha", "plodina", "plomba", "plout", "pluk", "plyn", "pobavit", "pobyt", "pochod", "pocit", "poctivec", "podat", "podcenit", "podepsat", "podhled", "podivit", "podklad", "podmanit", "podnik", "podoba", "podpora", "podraz", "podstata", "podvod", "podzim", "poezie", "pohanka", "pohnutka", "pohovor", "pohroma", "pohyb", "pointa", "pojistka", "pojmout", "pokazit", "pokles", "pokoj", "pokrok", "pokuta", "pokyn", "poledne", "polibek", "polknout", "poloha", "polynom", "pomalu", "pominout", "pomlka", "pomoc", "pomsta", "pomyslet", "ponechat", "ponorka", "ponurost", "popadat", "popel", "popisek", "poplach", "poprosit", "popsat", "popud", "poradce", "porce", "porod", "porucha", "poryv", "posadit", "posed", "posila", "poskok", "poslanec", "posoudit", "pospolu", "postava", "posudek", "posyp", "potah", "potkan", "potlesk", "potomek", "potrava", "potupa", "potvora", "poukaz", "pouto", "pouzdro", "povaha", "povidla", "povlak", "povoz", "povrch", "povstat", "povyk", "povzdech", "pozdrav", "pozemek", "poznatek", "pozor", "pozvat", "pracovat", "prahory", "praktika", "prales", "praotec", "praporek", "prase", "pravda", "princip", "prkno", "probudit", "procento", "prodej", "profese", "prohra", "projekt", "prolomit", "promile", "pronikat", "propad", "prorok", "prosba", "proton", "proutek", "provaz", "prskavka", "prsten", "prudkost", "prut", "prvek", "prvohory", "psanec", "psovod", "pstruh", "ptactvo", "puberta", "puch", "pudl", "pukavec", "puklina", "pukrle", "pult", "pumpa", "punc", "pupen", "pusa", "pusinka", "pustina", "putovat", "putyka", "pyramida", "pysk", "pytel", "racek", "rachot", "radiace", "radnice", "radon", "raft", "ragby", "raketa", "rakovina", "rameno", "rampouch", "rande", "rarach", "rarita", "rasovna", "rastr", "ratolest", "razance", "razidlo", "reagovat", "reakce", "recept", "redaktor", "referent", "reflex", "rejnok", "reklama", "rekord", "rekrut", "rektor", "reputace", "revize", "revma", "revolver", "rezerva", "riskovat", "riziko", "robotika", "rodokmen", "rohovka", "rokle", "rokoko", "romaneto", "ropovod", "ropucha", "rorejs", "rosol", "rostlina", "rotmistr", "rotoped", "rotunda", "roubenka", "roucho", "roup", "roura", "rovina", "rovnice", "rozbor", "rozchod", "rozdat", "rozeznat", "rozhodce", "rozinka", "rozjezd", "rozkaz", "rozloha", "rozmar", "rozpad", "rozruch", "rozsah", "roztok", "rozum", "rozvod", "rubrika", "ruchadlo", "rukavice", "rukopis", "ryba", "rybolov", "rychlost", "rydlo", "rypadlo", "rytina", "ryzost", "sadista", "sahat", "sako", "samec", "samizdat", "samota", "sanitka", "sardinka", "sasanka", "satelit", "sazba", "sazenice", "sbor", "schovat", "sebranka", "secese", "sedadlo", "sediment", "sedlo", "sehnat", "sejmout", "sekera", "sekta", "sekunda", "sekvoje", "semeno", "seno", "servis", "sesadit", "seshora", "seskok", "seslat", "sestra", "sesuv", "sesypat", "setba", "setina", "setkat", "setnout", "setrvat", "sever", "seznam", "shoda", "shrnout", "sifon", "silnice", "sirka", "sirotek", "sirup", "situace", "skafandr", "skalisko", "skanzen", "skaut", "skeptik", "skica", "skladba", "sklenice", "sklo", "skluz", "skoba", "skokan", "skoro", "skripta", "skrz", "skupina", "skvost", "skvrna", "slabika", "sladidlo", "slanina", "slast", "slavnost", "sledovat", "slepec", "sleva", "slezina", "slib", "slina", "sliznice", "slon", "sloupek", "slovo", "sluch", "sluha", "slunce", "slupka", "slza", "smaragd", "smetana", "smilstvo", "smlouva", "smog", "smrad", "smrk", "smrtka", "smutek", "smysl", "snad", "snaha", "snob", "sobota", "socha", "sodovka", "sokol", "sopka", "sotva", "souboj", "soucit", "soudce", "souhlas", "soulad", "soumrak", "souprava", "soused", "soutok", "souviset", "spalovna", "spasitel", "spis", "splav", "spodek", "spojenec", "spolu", "sponzor", "spornost", "spousta", "sprcha", "spustit", "sranda", "sraz", "srdce", "srna", "srnec", "srovnat", "srpen", "srst", "srub", "stanice", "starosta", "statika", "stavba", "stehno", "stezka", "stodola", "stolek", "stopa", "storno", "stoupat", "strach", "stres", "strhnout", "strom", "struna", "studna", "stupnice", "stvol", "styk", "subjekt", "subtropy", "suchar", "sudost", "sukno", "sundat", "sunout", "surikata", "surovina", "svah", "svalstvo", "svetr", "svatba", "svazek", "svisle", "svitek", "svoboda", "svodidlo", "svorka", "svrab", "sykavka", "sykot", "synek", "synovec", "sypat", "sypkost", "syrovost", "sysel", "sytost", "tabletka", "tabule", "tahoun", "tajemno", "tajfun", "tajga", "tajit", "tajnost", "taktika", "tamhle", "tampon", "tancovat", "tanec", "tanker", "tapeta", "tavenina", "tazatel", "technika", "tehdy", "tekutina", "telefon", "temnota", "tendence", "tenista", "tenor", "teplota", "tepna", "teprve", "terapie", "termoska", "textil", "ticho", "tiskopis", "titulek", "tkadlec", "tkanina", "tlapka", "tleskat", "tlukot", "tlupa", "tmel", "toaleta", "topinka", "topol", "torzo", "touha", "toulec", "tradice", "traktor", "tramp", "trasa", "traverza", "trefit", "trest", "trezor", "trhavina", "trhlina", "trochu", "trojice", "troska", "trouba", "trpce", "trpitel", "trpkost", "trubec", "truchlit", "truhlice", "trus", "trvat", "tudy", "tuhnout", "tuhost", "tundra", "turista", "turnaj", "tuzemsko", "tvaroh", "tvorba", "tvrdost", "tvrz", "tygr", "tykev", "ubohost", "uboze", "ubrat", "ubrousek", "ubrus", "ubytovna", "ucho", "uctivost", "udivit", "uhradit", "ujednat", "ujistit", "ujmout", "ukazatel", "uklidnit", "uklonit", "ukotvit", "ukrojit", "ulice", "ulita", "ulovit", "umyvadlo", "unavit", "uniforma", "uniknout", "upadnout", "uplatnit", "uplynout", "upoutat", "upravit", "uran", "urazit", "usednout", "usilovat", "usmrtit", "usnadnit", "usnout", "usoudit", "ustlat", "ustrnout", "utahovat", "utkat", "utlumit", "utonout", "utopenec", "utrousit", "uvalit", "uvolnit", "uvozovka", "uzdravit", "uzel", "uzenina", "uzlina", "uznat", "vagon", "valcha", "valoun", "vana", "vandal", "vanilka", "varan", "varhany", "varovat", "vcelku", "vchod", "vdova", "vedro", "vegetace", "vejce", "velbloud", "veletrh", "velitel", "velmoc", "velryba", "venkov", "veranda", "verze", "veselka", "veskrze", "vesnice", "vespodu", "vesta", "veterina", "veverka", "vibrace", "vichr", "videohra", "vidina", "vidle", "vila", "vinice", "viset", "vitalita", "vize", "vizitka", "vjezd", "vklad", "vkus", "vlajka", "vlak", "vlasec", "vlevo", "vlhkost", "vliv", "vlnovka", "vloupat", "vnucovat", "vnuk", "voda", "vodivost", "vodoznak", "vodstvo", "vojensky", "vojna", "vojsko", "volant", "volba", "volit", "volno", "voskovka", "vozidlo", "vozovna", "vpravo", "vrabec", "vracet", "vrah", "vrata", "vrba", "vrcholek", "vrhat", "vrstva", "vrtule", "vsadit", "vstoupit", "vstup", "vtip", "vybavit", "vybrat", "vychovat", "vydat", "vydra", "vyfotit", "vyhledat", "vyhnout", "vyhodit", "vyhradit", "vyhubit", "vyjasnit", "vyjet", "vyjmout", "vyklopit", "vykonat", "vylekat", "vymazat", "vymezit", "vymizet", "vymyslet", "vynechat", "vynikat", "vynutit", "vypadat", "vyplatit", "vypravit", "vypustit", "vyrazit", "vyrovnat", "vyrvat", "vyslovit", "vysoko", "vystavit", "vysunout", "vysypat", "vytasit", "vytesat", "vytratit", "vyvinout", "vyvolat", "vyvrhel", "vyzdobit", "vyznat", "vzadu", "vzbudit", "vzchopit", "vzdor", "vzduch", "vzdychat", "vzestup", "vzhledem", "vzkaz", "vzlykat", "vznik", "vzorek", "vzpoura", "vztah", "vztek", "xylofon", "zabrat", "zabydlet", "zachovat", "zadarmo", "zadusit", "zafoukat", "zahltit", "zahodit", "zahrada", "zahynout", "zajatec", "zajet", "zajistit", "zaklepat", "zakoupit", "zalepit", "zamezit", "zamotat", "zamyslet", "zanechat", "zanikat", "zaplatit", "zapojit", "zapsat", "zarazit", "zastavit", "zasunout", "zatajit", "zatemnit", "zatknout", "zaujmout", "zavalit", "zavelet", "zavinit", "zavolat", "zavrtat", "zazvonit", "zbavit", "zbrusu", "zbudovat", "zbytek", "zdaleka", "zdarma", "zdatnost", "zdivo", "zdobit", "zdroj", "zdvih", "zdymadlo", "zelenina", "zeman", "zemina", "zeptat", "zezadu", "zezdola", "zhatit", "zhltnout", "zhluboka", "zhotovit", "zhruba", "zima", "zimnice", "zjemnit", "zklamat", "zkoumat", "zkratka", "zkumavka", "zlato", "zlehka", "zloba", "zlom", "zlost", "zlozvyk", "zmapovat", "zmar", "zmatek", "zmije", "zmizet", "zmocnit", "zmodrat", "zmrzlina", "zmutovat", "znak", "znalost", "znamenat", "znovu", "zobrazit", "zotavit", "zoubek", "zoufale", "zplodit", "zpomalit", "zprava", "zprostit", "zprudka", "zprvu", "zrada", "zranit", "zrcadlo", "zrnitost", "zrno", "zrovna", "zrychlit", "zrzavost", "zticha", "ztratit", "zubovina", "zubr", "zvednout", "zvenku", "zvesela", "zvon", "zvrat", "zvukovod", "zvyk"}

var english_slip39 = []string{"academic", "acid", "acne", "acquire", "acrobat", "activity", "actress", "adapt", "adequate", "adjust", "admit", "adorn", "adult", "advance", "advocate", "afraid", "again", "agency", "agree", "aide", "aircraft", "airline", "airport", "ajar", "alarm", "album", "alcohol", "alien", "alive", "alpha", "already", "alto", "aluminum", "always", "amazing", "ambition", "amount", "amuse", "analysis", "anatomy", "ancestor", "ancient", "angel", "angry", "animal", "answer", "antenna", "anxiety", "apart", "aquatic", "arcade", "arena", "argue", "armed", "artist", "artwork", "aspect", "auction", "august", "aunt", "average", "aviation", "avoid", "award", "away", "axis", "axle", "beam", "beard", "beaver", "become", "bedroom", "behavior", "being", "believe", "belong", "benefit", "best", "beyond", "bike", "biology", "birthday", "bishop", "black", "blanket", "blessing", "blimp", "blind", "blue", "body", "bolt", "boring", "born", "both", "boundary", "bracelet", "branch", "brave", "breathe", "briefing", "broken", "brother", "browser", "bucket", "budget", "building", "bulb", "bulge", "bumpy", "bundle", "burden", "burning", "busy", "buyer", "cage", "calcium", "camera", "campus", "canyon", "capacity", "capital", "capture", "carbon", "cards", "careful", "cargo", "carpet", "carve", "category", "cause", "ceiling", "center", "ceramic", "champion", "change", "charity", "check", "chemical", "chest", "chew", "chubby", "cinema", "civil", "class", "clay", "cleanup", "client", "climate", "clinic", "clock", "clogs", "closet", "clothes", "club", "cluster", "coal", "coastal", "coding", "column", "company", "corner", "costume", "counter", "course", "cover", "cowboy", "cradle", "craft", "crazy", "credit", "cricket", "criminal", "crisis", "critical", "crowd", "crucial", "crunch", "crush", "crystal", "cubic", "cultural", "curious", "curly", "custody", "cylinder", "daisy", "damage", "dance", "darkness", "database", "daughter", "deadline", "deal", "debris", "debut", "decent", "decision", "declare", "decorate", "decrease", "deliver", "demand", "density", "deny", "depart", "depend", "depict", "deploy", "describe", "desert", "desire", "desktop", "destroy", "detailed", "detect", "device", "devote", "diagnose", "dictate", "diet", "dilemma", "diminish", "dining", "diploma", "disaster", "discuss", "disease", "dish", "dismiss", "display", "distance", "dive", "divorce", "document", "domain", "domestic", "dominant", "dough", "downtown", "dragon", "dramatic", "dream", "dress", "drift", "drink", "drove", "drug", "dryer", "duckling", "duke", "duration", "dwarf", "dynamic", "early", "earth", "easel", "easy", "echo", "eclipse", "ecology", "edge", "editor", "educate", "either", "elbow", "elder", "election", "elegant", "element", "elephant", "elevator", "elite", "else", "email", "emerald", "emission", "emperor", "emphasis", "employer", "empty", "ending", "endless", "endorse", "enemy", "energy", "enforce", "engage", "enjoy", "enlarge", "entrance", "envelope", "envy", "epidemic", "episode", "equation", "equip", "eraser", "erode", "escape", "estate", "estimate", "evaluate", "evening", "evidence", "evil", "evoke", "exact", "example", "exceed", "exchange", "exclude", "excuse", "execute", "exercise", "exhaust", "exotic", "expand", "expect", "explain", "express", "extend", "extra", "eyebrow", "facility", "fact", "failure", "faint", "fake", "false", "family", "famous", "fancy", "fangs", "fantasy", "fatal", "fatigue", "favorite", "fawn", "fiber", "fiction", "filter", "finance", "findings", "finger", "firefly", "firm", "fiscal", "fishing", "fitness", "flame", "flash", "flavor", "flea", "flexible", "flip", "float", "floral", "fluff", "focus", "forbid", "force", "forecast", "forget", "formal", "fortune", "forward", "founder", "fraction", "fragment", "frequent", "freshman", "friar", "fridge", "friendly", "frost", "froth", "frozen", "fumes", "funding", "furl", "fused", "galaxy", "game", "garbage", "garden", "garlic", "gasoline", "gather", "general", "genius", "genre", "genuine", "geology", "gesture", "glad", "glance", "glasses", "glen", "glimpse", "goat", "golden", "graduate", "grant", "grasp", "gravity", "gray", "greatest", "grief", "grill", "grin", "grocery", "gross", "group", "grownup", "grumpy", "guard", "guest", "guilt", "guitar", "gums", "hairy", "hamster", "hand", "hanger", "harvest", "have", "havoc", "hawk", "hazard", "headset", "health", "hearing", "heat", "helpful", "herald", "herd", "hesitate", "hobo", "holiday", "holy", "home", "hormone", "hospital", "hour", "huge", "human", "humidity", "hunting", "husband", "hush", "husky", "hybrid", "idea", "identify", "idle", "image", "impact", "imply", "improve", "impulse", "include", "income", "increase", "index", "indicate", "industry", "infant", "inform", "inherit", "injury", "inmate", "insect", "inside", "install", "intend", "intimate", "invasion", "involve", "iris", "island", "isolate", "item", "ivory", "jacket", "jerky", "jewelry", "join", "judicial", "juice", "jump", "junction", "junior", "junk", "jury", "justice", "kernel", "keyboard", "kidney", "kind", "kitchen", "knife", "knit", "laden", "ladle", "ladybug", "lair", "lamp", "language", "large", "laser", "laundry", "lawsuit", "leader", "leaf", "learn", "leaves", "lecture", "legal", "legend", "legs", "lend", "length", "level", "liberty", "library", "license", "lift", "likely", "lilac", "lily", "lips", "liquid", "listen", "literary", "living", "lizard", "loan", "lobe", "location", "losing", "loud", "loyalty", "luck", "lunar", "lunch", "lungs", "luxury", "lying", "lyrics", "machine", "magazine", "maiden", "mailman", "main", "makeup", "making", "mama", "manager", "mandate", "mansion", "manual", "marathon", "march", "market", "marvel", "mason", "material", "math", "maximum", "mayor", "meaning", "medal", "medical", "member", "memory", "mental", "merchant", "merit", "method", "metric", "midst", "mild", "military", "mineral", "minister", "miracle", "mixed", "mixture", "mobile", "modern", "modify", "moisture", "moment", "morning", "mortgage", "mother", "mountain", "mouse", "move", "much", "mule", "multiple", "muscle", "museum", "music", "mustang", "nail", "national", "necklace", "negative", "nervous", "network", "news", "nuclear", "numb", "numerous", "nylon", "oasis", "obesity", "object", "observe", "obtain", "ocean", "often", "olympic", "omit", "oral", "orange", "orbit", "order", "ordinary", "organize", "ounce", "oven", "overall", "owner", "paces", "pacific", "package", "paid", "painting", "pajamas", "pancake", "pants", "papa", "paper", "parcel", "parking", "party", "patent", "patrol", "payment", "payroll", "peaceful", "peanut", "peasant", "pecan", "penalty", "pencil", "percent", "perfect", "permit", "petition", "phantom", "pharmacy", "photo", "phrase", "physics", "pickup", "picture", "piece", "pile", "pink", "pipeline", "pistol", "pitch", "plains", "plan", "plastic", "platform", "playoff", "pleasure", "plot", "plunge", "practice", "prayer", "preach", "predator", "pregnant", "premium", "prepare", "presence", "prevent", "priest", "primary", "priority", "prisoner", "privacy", "prize", "problem", "process", "profile", "program", "promise", "prospect", "provide", "prune", "public", "pulse", "pumps", "punish", "puny", "pupal", "purchase", "purple", "python", "quantity", "quarter", "quick", "quiet", "race", "racism", "radar", "railroad", "rainbow", "raisin", "random", "ranked", "rapids", "raspy", "reaction", "realize", "rebound", "rebuild", "recall", "receiver", "recover", "regret", "regular", "reject", "relate", "remember", "remind", "remove", "render", "repair", "repeat", "replace", "require", "rescue", "research", "resident", "response", "result", "retailer", "retreat", "reunion", "revenue", "review", "reward", "rhyme", "rhythm", "rich", "rival", "river", "robin", "rocky", "romantic", "romp", "roster", "round", "royal", "ruin", "ruler", "rumor", "sack", "safari", "salary", "salon", "salt", "satisfy", "satoshi", "saver", "says", "scandal", "scared", "scatter", "scene", "scholar", "science", "scout", "scramble", "screw", "script", "scroll", "seafood", "season", "secret", "security", "segment", "senior", "shadow", "shaft", "shame", "shaped", "sharp", "shelter", "sheriff", "short", "should", "shrimp", "sidewalk", "silent", "silver", "similar", "simple", "single", "sister", "skin", "skunk", "slap", "slavery", "sled", "slice", "slim", "slow", "slush", "smart", "smear", "smell", "smirk", "smith", "smoking", "smug", "snake", "snapshot", "sniff", "society", "software", "soldier", "solution", "soul", "source", "space", "spark", "speak", "species", "spelling", "spend", "spew", "spider", "spill", "spine", "spirit", "spit", "spray", "sprinkle", "square", "squeeze", "stadium", "staff", "standard", "starting", "station", "stay", "steady", "step", "stick", "stilt", "story", "strategy", "strike", "style", "subject", "submit", "sugar", "suitable", "sunlight", "superior", "surface", "surprise", "survive", "sweater", "swimming", "swing", "switch", "symbolic", "sympathy", "syndrome", "system", "tackle", "tactics", "tadpole", "talent", "task", "taste", "taught", "taxi", "teacher", "teammate", "teaspoon", "temple", "tenant", "tendency", "tension", "terminal", "testify", "texture", "thank", "that", "theater", "theory", "therapy", "thorn", "threaten", "thumb", "thunder", "ticket", "tidy", "timber", "timely", "ting", "tofu", "together", "tolerate", "total", "toxic", "tracks", "traffic", "training", "transfer", "trash", "traveler", "treat", "trend", "trial", "tricycle", "trip", "triumph", "trouble", "true", "trust", "twice", "twin", "type", "typical", "ugly", "ultimate", "umbrella", "uncover", "undergo", "unfair", "unfold", "unhappy", "union", "universe", "unkind", "unknown", "unusual", "unwrap", "upgrade", "upstairs", "username", "usher", "usual", "valid", "valuable", "vampire", "vanish", "various", "vegan", "velvet", "venture", "verdict", "verify", "very", "veteran", "vexed", "victim", "video", "view", "vintage", "violence", "viral", "visitor", "visual", "vitamins", "vocal", "voice", "volume", "voter", "voting", "walnut", "warmth", "warn", "watch", "wavy", "wealthy", "weapon", "webcam", "welcome", "welfare", "western", "width", "wildlife", "window", "wine", "wireless", "wisdom", "withdraw", "wits", "wolf", "woman", "work", "worthy", "wrap", "wrist", "writing", "wrote", "year", "yelp", "yield", "yoga", "zero"}

func selectWordList() {

	printStyled("\n{cyan}{bold}Select a word list:\n\n")
	printStyled("{cyan}1. {bold}English (slip39, 1024 words used by Trezor)\n")
	printStyled("{cyan}2. {bold}English (BIP39, 2048 words)\n")
	printStyled("{cyan}3. {bold}Czech (BIP39, 2048 words)\n")
	reader := bufio.NewReader(os.Stdin)
	var selection int
	for {
		printStyled("\n{cyan}Enter the number of the word list to use: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		selection, err = strconv.Atoi(input)
		if err == nil && selection >= 1 && selection <= 3 {
			break
		}
		printStyled("{red}Invalid selection. Please try again.\n")
	}
	switch selection {
	case 1:
		words = english_slip39
	case 2:
		words = english_bip39
	case 3:
		words = cezch_bip39
	}
	printStyled("\n{green}Word list successfully loaded!\n\n")
}

func findWord(word string) int {
	for i, w := range words {
		if w == word {
			return i
		}
	}
	return -1
}

func suggestWords(word string) []string {
	type suggestion struct {
		word  string
		score int
	}
	var suggestions []suggestion

	for _, w := range words {
		distance := levenshtein.ComputeDistance(word, w)
		suggestions = append(suggestions, suggestion{word: w, score: distance})
	}
	sort.Slice(suggestions, func(i, j int) bool {
		return suggestions[i].score < suggestions[j].score
	})
	result := []string{}
	for i, s := range suggestions {
		if i >= 3 {
			break
		}
		result = append(result, s.word)
	}
	return result
}

func isWeakPassword(password string) bool {
	if len(password) < 8 {
		return true
	}
	hasLower := strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz")
	hasUpper := strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	hasNumber := strings.ContainsAny(password, "0123456789")
	hasSpecial := strings.ContainsAny(password, "!@#$%^&*()-_=+[]{}|;:',.<>?/")

	return !(hasLower && hasUpper && hasNumber && hasSpecial)
}

func hashRepeatedly(data []byte, iterations int) []byte {
	hash := data
	for i := 0; i < iterations; i++ {
		digest := sha3.Sum512(hash)
		hash = digest[:]
	}
	return hash
}

func bytesToBitString(data []byte) string {
	var bitString strings.Builder
	for _, b := range data {
		for i := 7; i >= 0; i-- {
			if (b & (1 << i)) != 0 {
				bitString.WriteByte('1')
			} else {
				bitString.WriteByte('0')
			}
		}
	}
	return bitString.String()
}

func splitString(input string, length int) []string {
	if length <= 0 {
		return []string{}
	}
	var result []string
	for i := 0; i < len(input); i += length {
		end := i + length
		if end > len(input) {
			end = len(input)
		}
		result = append(result, input[i:end])
	}
	return result
}

func xorBitStrings(bits1, bits2 string) string {
	if len(bits1) < len(bits2) {
		bits1 = strings.Repeat("0", len(bits2)-len(bits1)) + bits1
	} else if len(bits2) < len(bits1) {
		bits2 = strings.Repeat("0", len(bits1)-len(bits2)) + bits2
	}
	var result strings.Builder
	for i := 0; i < len(bits1); i++ {
		if bits1[i] == bits2[i] {
			result.WriteByte('0')
		} else {
			result.WriteByte('1')
		}
	}
	return result.String()
}

func bitsToInt(bits string) int {
	result := new(big.Int)
	result.SetString(bits, 2)
	return int(result.Int64())
}

func intToBits(value int, bitLength int) string {
	bitString := strconv.FormatInt(int64(value), 2)
	if len(bitString) < bitLength {
		padding := bitLength - len(bitString)
		bitString = strings.Repeat("0", padding) + bitString
	}
	return bitString
}

func printStyled(text string) {
	reset := "\033[0m"
	styles := map[string]string{
		"red":       "\033[31m",
		"green":     "\033[32m",
		"yellow":    "\033[33m",
		"cyan":      "\033[36m",
		"bold":      "\033[1m",
		"underline": "\033[4m",
	}
	for style, code := range styles {
		placeholder := fmt.Sprintf("{%s}", style)
		text = strings.ReplaceAll(text, placeholder, code)
	}
	text = strings.ReplaceAll(text, "{reset}", reset)
	fmt.Print(text + reset)
}

func pressAnyKey() {
	printStyled("\n{bold}{cyan}Press any key to continue...\n")
	bufio.NewReader(os.Stdin).ReadByte()
}

func choice(message string, first string, second string, letter1 string, letter2 string) bool {
	var userInput string
	letter1 = strings.ToUpper(letter1)
	letter2 = strings.ToUpper(letter2)
	reader := bufio.NewReader(os.Stdin)
	for {
		printStyled("{bold}{cyan}" + message)
		printStyled("\nPlease Choose {bold}{cyan}(" + letter1 + ") {reset}" + first + ", or {bold}{cyan}(" + letter2 + ") {reset}" + second + ": ")
		userInput, _ = reader.ReadString('\n')
		userInput = strings.ToUpper(strings.TrimSpace(userInput))
		if userInput == letter1 || userInput == letter2 {
			break
		}
		printStyled("\n{red}Invalid choice!\n")
	}
	return userInput == letter1
}
func getPassword(recover bool) string {
	reader := bufio.NewReader(os.Stdin)
	var password1, password2 string
	for {
		printStyled("\n{cyan}Enter password: ")
		password1, _ = reader.ReadString('\n')
		password1 = strings.TrimSpace(password1)

		printStyled("{cyan}Confirm the password: ")
		password2, _ = reader.ReadString('\n')
		password2 = strings.TrimSpace(password2)

		if password1 != password2 {
			printStyled("{red}{bold}\nError: Passwords do not match. Try again.")
			continue
		}

		if !recover && isWeakPassword(password1) {
			printStyled("\n{yellow}Warning: Your password is weak. It should be at least 8 characters long\n")
			printStyled("{yellow}and include a mix of uppercase, lowercase, numbers, and special characters.\n")
			printStyled("{yellow}Type {red}'YES'{yellow} if you want to continue with this password: ")
			confirmation, _ := reader.ReadString('\n')
			confirmation = strings.TrimSpace(confirmation)
			if confirmation == "YES" {
				printStyled("\n{green}Ok, weak password accepted.")
				break
			} else {
				printStyled("\n{green}Please enter a stronger password.")
				continue
			}
		} else {
			printStyled("\n{green}Password accepted.")
			break
		}
	}

	if !recover {
		printStyled("\n\n{yellow}Don't forget your password - there is {underline}NO WAY{reset}{yellow} to recover it!\n\n")
	}
	return password1
}
func main() {

	printStyled("\n\n{cyan}{bold}{underline}Welcome to the wallet word scrambler\n\n")
	printStyled("A password will be used to scramble your backup words\n")
	printStyled("In order for the program to work a wordlist must be present at the current folder\n")
	printStyled("The wordlist must match your wallet backup\n\n")
	printStyled("{red}Warning:\n")
	printStyled("{yellow}This program is meant to run on a fresh formatted and air-gapped machine\n")
	printStyled("{yellow}It is not safe to run it on a machine connected to any kind of network\n")
	printStyled("{yellow}{bold}SECURE WIPE{reset}{yellow} your machine after use\n\n")

	pressAnyKey()

	selectWordList()

	recover := choice("Do you want to recover a wallet or create (scramble) a new one?", "Recover", "Create", "R", "C")
	reader := bufio.NewReader(os.Stdin)
	var password = getPassword(recover)
	printStyled("\n{cyan}Gererating a key from your password.\nThis is an intentionaly slow process, please wait.\n\n")
	argon2Hash := hashRepeatedly([]byte("this is just something to season the dish"), 4847868)
	memory := uint32(1024 * 1024)
	time := uint32(16)
	threads := uint8(1)
	keyLen := uint32(64)
	rounds := 10
	progress := 0
	for i := 1; i <= rounds; i++ {
		argon2Hash = hashRepeatedly(argon2Hash, i)
		argon2Hash = argon2.IDKey([]byte(password), argon2Hash, time, memory, threads, keyLen)
		progress = (i * 100) / rounds
		showProgress := ".." + strconv.Itoa(progress) + "%"
		printStyled(showProgress)
	}

	keyBits := bytesToBitString(argon2Hash)
	wordBitSize := int(math.Log2(float64(len(words))))
	keyBitsWords := splitString(keyBits, wordBitSize)

	printStyled("\n\n{green}Key generated.\n")

	var walletWordCount int
	for {
		printStyled("\n{cyan}Enter the number of words in your wallet (12-33): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		walletWordCount, err = strconv.Atoi(input)
		if err == nil && walletWordCount >= 12 && walletWordCount <= 33 {
			break
		}
		fmt.Println("Invalid input. Please enter a number between 12 and 33.")
	}
	newWords := make([]string, walletWordCount)
	for i := 0; i < walletWordCount; i++ {
		var word string
		for {
			fmt.Printf("Enter word %d: ", i+1)
			word, _ = reader.ReadString('\n')
			word = strings.TrimSpace(word)
			wordIndex := findWord(word)
			if wordIndex == -1 {
				suggestions := suggestWords(word)
				printStyled("{red}Invalid word. The word must exist in the wordlist.\n")
				if len(suggestions) > 0 {
					printStyled("{yellow}Did you mean:\n")
					for _, suggestion := range suggestions {
						printStyled("{white}" + suggestion + "\n")
					}
				}
			} else {
				wordBits := intToBits(wordIndex, wordBitSize)
				xorResult := xorBitStrings(keyBitsWords[i], wordBits)
				newWordIndex := bitsToInt(xorResult)
				newWords[i] = words[newWordIndex]
				break
			}
		}
	}
	if !recover {

		printStyled("\n{bold}{underline}{cyan}Here are your new wallet words\n\n")
	} else {
		printStyled("\n{bold}{underline}{cyan}Here are your recovered wallet words\n\n")
	}
	numberPadding := len(fmt.Sprintf("%d", len(newWords)))
	for i, word := range newWords {
		fmt.Printf("%*d. %s\n", numberPadding, i+1, word)
	}

	if !recover {
		printStyled("\n\n{yellow}Write the words down and store them in a safe place.\n\n")
	}
	pressAnyKey()
}
