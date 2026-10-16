# go-reloaded

Outil de complétion, d'édition et de correction automatique de texte, écrit en Go.

Le programme lit un fichier texte, applique une série de modifications indiquées par des marqueurs placés dans le texte, puis écrit le résultat dans un nouveau fichier.

## Utilisation

```bash
go run . <fichier_entree> <fichier_sortie>
```

Exemple :

```bash
$ cat sample.txt
It has been 10 (bin) years and 1E (hex) files were added

$ go run . sample.txt result.txt

$ cat result.txt
It has been 2 years and 30 files were added
```

Si le nombre d'arguments est incorrect ou si le fichier d'entrée ne peut pas être lu, le programme affiche un message d'erreur et s'arrête.

## Prérequis

- Go 1.21 ou supérieur
- Uniquement des packages de la bibliothèque standard (`fmt`, `os`, `strconv`, `strings`, `unicode`)

## Règles implémentées

| Marqueur | Effet | Exemple | Résultat |
|---|---|---|---|
| `(hex)` | Convertit le mot précédent de l'hexadécimal vers le décimal | `1E (hex)` | `30` |
| `(bin)` | Convertit le mot précédent du binaire vers le décimal | `10 (bin)` | `2` |
| `(up)` | Met le mot précédent en majuscules | `go (up)` | `GO` |
| `(low)` | Met le mot précédent en minuscules | `SHOUTING (low)` | `shouting` |
| `(cap)` | Met une majuscule au début du mot précédent | `bridge (cap)` | `Bridge` |

Les marqueurs `(up)`, `(low)` et `(cap)` acceptent aussi un nombre pour s'appliquer à plusieurs mots :

| Marqueur | Exemple | Résultat |
|---|---|---|
| `(up, n)` | `so exciting (up, 2)` | `SO EXCITING` |
| `(low, n)` | `STOP THIS NOW (low, 3)` | `stop this now` |
| `(cap, n)` | `eiffel tower (cap, 2)` | `Eiffel Tower` |

Dans tous les cas, le marqueur est supprimé du texte final.

## Fonctionnement

1. Le fichier d'entrée est lu avec `os.ReadFile`.
2. Le texte est découpé en mots avec `strings.Fields`.
3. Chaque règle est appliquée l'une après l'autre sur la liste de mots, dans cet ordre :
   `convertHex`, `convertBin`, `convertUp`, `convertLow`, `convertCap`.
4. Les mots sont recollés avec un espace (`strings.Join`) et écrits dans le fichier de sortie avec `os.WriteFile`.

### Détail des fonctions

- **`convertHex` / `convertBin`** : repèrent le marqueur, convertissent le mot précédent avec `strconv.ParseInt` (base 16 ou base 2), puis le réécrivent en base 10 avec `strconv.FormatInt`. Si le mot n'est pas un nombre valide, il est laissé tel quel.
- **`convertUp` / `convertLow` / `convertCap`** : gèrent deux cas.
  - Le marqueur simple, par exemple `(up)`, qui modifie un seul mot.
  - Le marqueur avec nombre, par exemple `(up, 2)`. `strings.Fields` le découpe en deux éléments (`(up,` et `2)`), donc le nombre est lu dans l'élément suivant avec `strings.TrimSuffix` et `strconv.Atoi`. La modification s'applique ensuite aux `n` mots précédents, sans dépasser le début du texte.
- **`capitalize`** : met tout le mot en minuscules, puis la première lettre en majuscule avec `unicode.ToUpper`. Le mot est manipulé sous forme de `[]rune` pour gérer correctement les caractères accentués.

Pour supprimer un marqueur de la liste, le programme utilise `append(words[:i], words[i+1:]...)`, puis recule l'index (`i--`) pour ne pas sauter le mot suivant.

## Structure du projet

```
go-reloaded/
├── go.mod
├── main.go
└── README.md
```
