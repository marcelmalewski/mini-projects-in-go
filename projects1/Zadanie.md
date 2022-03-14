Paulina i Marcel muszą przeprowadzić pewne operacje na pewnym pliku. Zależy od tego ich życie. Alternatywą jest defenestracja. 

Plik z którego mają do wczytania dane pochodzi od Doktora Wiesława Pawłowskiego,
chociaż to Szanowny Wielki Polak Sienkiewicz jest oryginalnym autorem. Plik nazywa się 'ogniem_i_mieczem.txt' i jest gdzieś tu do pobrania.

W pliku znajdują się słowa w zdaniach w akapitach. Nie jedno - jest ich więcej - jednak nie jest ich też zbyt wiele. Tym nie powinni się martwić.

Pierwszym krokiem jest pobranie 
[Brawl Stars](https://play.google.com/store/apps/details?id=com.supercell.brawlstars).
Następnym krokiem jest przystąpienie do realizacji zadania.

Twoim zadaniem jest pomóć Paulinie i Marcelowi przeżyć.

--- 

Paulina i Marcel muszą policzyć ile razy każde słowo występuje w tekście. Słowa różniące się tylko wielkością liter uznawane są za te same słowo. Różne odmiany tego samego słowa uznawane są za różne słowa. 

Przykład odnośnie interpunkcji:
```
Dniepru, a od Dniestru - niedaleko za Humaniem, a potem... już hen, ku
```
należy interpretować jako słowa
```
dniepru a od dniestru niedaleko za humaniem a potem już hen ku
```


Tak zebrane pary, tj np.
```
śmierć - 5
życie - 2
```
należy posortować według ilości wystąpień.

Następnie, należy wyświetlić wyniki. Przy wyświetlaniu wyników nie wyświetlamy liczb, zamiast tego wyświetlamy ilość wystąpień wizualnie, kwadratami. Przykładowo:
```
śmireć - 🟩🟩🟩🟩🟩
życie - 🟩🟩
```

---
Paulina jest optymistą i nie uwzględnia w swoich wynikach żadnych słów zawierających podciąg `nie`.

Marcel natomiast nienawidzi języka polskiego, a przynajmniej polskich znaków diakrytycznych i tym samym nie uwzględnia w swoich wynikach żadnych słów zawierających jakikolwiek z tych znaków.

Pomóż Paulinie i Marcelowi uzyskać odpowiednie dla nich wyniki.

---
Wyniki wyświetl w okienku korzystając z dowolnej biblioteki do interfejsów graficznych.

---

## NOWOŚĆ!
Tak się składa, że zamiast tylko zielonych kwadratów, powinny być użyte następujące rodzaje kwadratów w zależności od ilości liter w słowie: 
🟫 - 1 litera
🟥 - 2 litery
🟧🟨🟩🟩🟦🟪 - analogicznie dla 3, 4, 5, 6, 7, 8 liter
⬜ - 9 liter i więcej.

