# Übungen

**Für jeden Übungsteil steht ein entsprechendes Verzeichnis zum Erstellen der Lösung zur Verfügung.**

## a) Variablen & Konstanten

### ✅Variablen

Deklarieren Sie auf verschiedene Weisen Variablen.

Welche Zuweisungen von einer zu anderen gehen, welche nicht?

### ✅ Konstanten

Deklarieren Sie nicht-gruppierte und gruppierte Konstanten.

Experimenten Sie mit untypisierten Konstanten (z.B. `const x = 1`) und welchen typisierten Variablen diese zugewiesen
werden kann.

### ✅ Iota

Schreiben Sie eine Gruppe an Konstanten, welche die Werte `KB`, `MB` und `GB` mit der jeweiligen Anzahl an
repräsentierten Bytes deklariert.

Nutzen Sie hierfür `iota` zur automatischen Berechnung.

Tipp: es gibt einen Bit-Shift Operator

## b) Pointers

### ✅ miniSort

Verändern Sie die Funktion `miniSort` und deren Aufruf derart, dass die gewünschte Ausgabe
erfolgt.

## ✅ c) Types

Wir wollen Zentimeter in Meter umwandeln können -- und dafür eigene Typen haben, wie z.B.

````go
type centimeter int
````

Deklarieren Sie eigene Typen für `centimeter` und `meter`. Wo muss die Deklaration erfolgen?

Schreiben Sie eine Funktion, die Zentimeter in Meter umrechnet.

Geben Sie das Ergebnis aus, sodass der gewünschte Output erfolgt.

## e) ✅ For Schleifen

Deklarieren Sie ein Array von Zahlen, z.B. `int` oder `float32`.

Schreiben Sie eine Funktion, welche das Array als Parameter annimmt und
die Summe aller Zahlen zurückgibt.

Wenn Sie ein `float32` Array nehmen, dann könnten Sie z.B. 10-mal die 0.1 addieren lassen. Was kommt da raus?

## f) ✅ If Abfragen

Schreiben Sie eine Funktion, welche für das Array von Zahlen (siehe oben)
nun sowohl den kleinsten als auch den größten Wert zurückgibt.

## g) ✅ Switch

Wir wollen wissen, ob heute Wochenende ist (okay, recht unwahrscheinlich während des Trainings).

Dazu bitte im StdLib Package "time" (https://pkg.go.dev/time) die Funktion `Now()` und `Weekday()` anschauen. Ebenso
gibt es dort Konstanten für die einzelnen Wochentage (welchen Typ haben diese?).

Schreiben Sie eine `switch` Anweisung, die für den heutigen Tag ausgibt, ob es gerade Wochenende ist oder nicht.

## h) ✅ Strings und Runes

### CamelCase ✅

Für einen beliebigen String wie z.B. "Heute lerne ich Go" soll die Camel-Case-Variante in der Form "HeuteLerneIchGo"
berechnet werden.

Schreiben Sie die dafür notwendige Funktion.

Hilfreich sind die StdLib Packages "strings" und "unicode".

### kebap-case / snake-case ✅

Ergänzend können Sie eine Transformation in Kebab-Case vornehmen: "heute-lerne-ich-go"

## i) Structs

Als Entwickler sind Sie mit einer JSON Schnittstelle konfrontiert, die etwas unschöne Attributnamen verwendet.

Mit Go möchten Sie nun (mittels eines Structs) das folgende JSON erzeugen:

````json
{
  "produktId": "P-123",
  "NAME": "Pizza Salami",
  "Preis": {
    "betrag-je-stueck": 12.99,
    "währung": "EUR"
  }
}
````

Erstellen Sie die dafür nötigen Structs. Diese sollen nur englische Attribute verwenden
(also "productId" statt "produktId") und in Camel-Case notiert werden ("unitPrice" statt "betrag-je-stueck").

Legen Sie dann eine Struct Instanz mit den obigen Daten an.

Exportieren Sie den Struct in JSON mit `json.Marshal(p1)` -- aber Achtung, es gibt noch ein Problem...

## j) Variadische Funktionen

Schreiben Sie eine Funktion, die alle übergebenen `int` Werte verdoppelt, sodass sich diese Werte in
der **aufrufenden** Funktion auch verändern. Anders gesagt: Jeder einzelne übergebene Wert soll am Ort 
seiner Deklaration verdoppelt worden sein.

Es soll kein Array übergeben werden, sondern einzelne Werte! Ebenso darf kein Wert zurückgegeben werden.

Welche Signatur (d.h. Parameterdeklarationen) muss diese Funktion haben?

## l) First Class Citizen Functions

Deklarieren Sie einen Funktionstyp mit folgender Signatur:

````go
func(int, int) int
````

Implementieren Sie diese Funktion in mehreren Varianten:

* Addition
* Subtraktion
* Multiplikation

Rufen Sie nun ihre Funktion (auch gerne verschachtelt) auf und lassen
das Ergebnis ausgeben.

## o) Slices

Implementieren Sie einen **Stack** von `string` Werten mithilfe eines Slice.

Der Stack soll über folgende Funktionen genutzt werden können:

* `push` -- hier wird ein Wert auf den Stack gelegt
* `pop` -- es wird ein Wert von oben vom Stack genommen und ausgeben
* `peek` -- das oberste Element wird ausgegeben ohne den Stack zu verändern

## p) Maps

Angenommen es gibt ein Struct mit den Attributen

* isbn (string)
* author (string)
* publisher (string)

Schreiben Sie eine Funktion, welche die Bücher je ISBN indiziert (je Key ein Buch).

Schreiben Sie dann zwei Funktionen, welche die Bücher in einer Map gruppiert (je Key mehrere Bücher):

* nach Autor (`groupBooksByAuthor()`)
* nach Verlag (`groupBooksByPublisher()`)

**Bonus:** Gruppieren Sie die Bücher erneut nach Autor und Verlag. Nutzen Sie jetzt aber eine jeweils per Closure
erstellte Funktion, die den Gruppierungsschlüssel (key) aus jedem Buch extrahiert und für die Gruppierung als Key nimmt.
Wir wollen damit erreichen, dass die (sehr hohe) Code-Duplizierung der `groupBy...()` Funktionen entfällt. Die
Key-Extractor Funktion übernimmt dann in der neuen `groupBy()` Funktion die Individualisierung, die bisher in den
zwei Ausprägungen existiert hat.