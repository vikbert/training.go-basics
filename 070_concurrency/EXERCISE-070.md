# Übungen zum Thema "Concurrency"

## a) Channels: Lesen und Schreiben

Starten Sie eine Goroutine, die in einen Channel einen Wert schreibt (z.B. ein `int`).

Dieser Wert soll in der `main()` Funktion ausgelesen und ausgegeben werden.

## b) Gemeinsamer Start

Lassen Sie drei Goroutinen starten, die eine zufällige Anzahl von 1 - 10 Sekunden brauchen, bis diese einsatzbereit
sind. Dazu können Sie die fertige Funktion `startup()` nutzen.

Sobald jede Goroutine bereit ist, soll sie auf ein gemeinsames Startsignal (mittels Channel) warten und dann
die aktuelle Zeit ausgeben.

Wie wissen wir, wann jede Goroutine bereit ist? Mögliche Optionen, von einfach nach kompliziert:

1. in `main()` einfach viele Sekunden warten
2. Oder man löst es mit einer `sync.WaitGroup`
3. Oder mit einem zusätzlichen Kanal

In jedem Fall schreiben wir dann in den Startsignal-Channel Werte, um jede Goroutine nahezu gleichzeitig
loslaufen zu lassen.

## e) Entweder Oder

Starten Sie eine Goroutine, die jede Sekunde einen `int` Wert in einen Kanal schreibt.

Starten Sie eine andere Goroutine, die alle zwei Sekunden einen `string` in einen anderen Kanal schreibt.

In der `main()` Methode nutzen Sie ein `select`, um entweder auf die eine oder auf die andere Nachricht zu reagieren.