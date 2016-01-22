package main

import "fmt"
import "levenshtein"

func main() {

    var elements []string
    elements = append(elements,"hydrogen")
    elements = append(elements,"helium")
    elements = append(elements,"lithium")
    elements = append(elements,"beryllium")
    elements = append(elements,"boron")
    elements = append(elements,"carbon")
    elements = append(elements,"nitrogen")
    elements = append(elements,"oxygen")
    elements = append(elements,"fluorine")
    elements = append(elements,"neon")
    elements = append(elements,"sodium")
    elements = append(elements,"magnesium")
    elements = append(elements,"aluminium")
    elements = append(elements,"silicon")
    elements = append(elements,"phosphorus")
    elements = append(elements,"sulfur")
    elements = append(elements,"chlorine")
    elements = append(elements,"argon")
    elements = append(elements,"potassium")
    elements = append(elements,"calcium")
    elements = append(elements,"scandium")
    elements = append(elements,"titanium")
    elements = append(elements,"vanadium")
    elements = append(elements,"chromium")
    elements = append(elements,"manganese")
    elements = append(elements,"iron")
    elements = append(elements,"cobalt")
    elements = append(elements,"nickel")
    elements = append(elements,"copper")
    elements = append(elements,"zinc")
    elements = append(elements,"gallium")
    elements = append(elements,"germanium")
    elements = append(elements,"arsenic")
    elements = append(elements,"selenium")
    elements = append(elements,"bromine")
    elements = append(elements,"krypton")
    elements = append(elements,"strontium")
    elements = append(elements,"rubidium")
    elements = append(elements,"yttrium")
    elements = append(elements,"zirconium")
    elements = append(elements,"niobium")
    elements = append(elements,"molybdenum")
    elements = append(elements,"technetium")
    elements = append(elements,"ruthenium");
    elements = append(elements,"rhodium");
    elements = append(elements,"palladium");
    elements = append(elements,"silver");
    elements = append(elements,"cadmium");
    elements = append(elements,"indium");
    elements = append(elements,"tin");
    elements = append(elements,"antimony");
    elements = append(elements,"tellurium");
    elements = append(elements,"iodine");
    elements = append(elements,"xenon");
    elements = append(elements,"caesium");
    elements = append(elements,"barium");
    elements = append(elements,"lanthanum");
    elements = append(elements,"cerium");
    elements = append(elements,"praseodymium");
    elements = append(elements,"neodymium");
    elements = append(elements,"promethium");
    elements = append(elements,"samarium");
    elements = append(elements,"europium");
    elements = append(elements,"gadolinium");
    elements = append(elements,"terbium");
    elements = append(elements,"dysprosium");
    elements = append(elements,"holmium");
    elements = append(elements,"erbium");
    elements = append(elements,"thulium");
    elements = append(elements,"ytterbium");
    elements = append(elements,"lutetium");
    elements = append(elements,"hafnium");
    elements = append(elements,"tantalum");
    elements = append(elements,"tungsten");
    elements = append(elements,"osmium");
    elements = append(elements,"rhenium");
    elements = append(elements,"iridium");
    elements = append(elements,"platinum");
    elements = append(elements,"gold");
    elements = append(elements,"mercury");
    elements = append(elements,"thallium");
    elements = append(elements,"lead");
    elements = append(elements,"bismuth");
    elements = append(elements,"polonium");
    elements = append(elements,"astatine");
    elements = append(elements,"radon");
    elements = append(elements,"francium");
    elements = append(elements,"radium");
    elements = append(elements,"actinium");
    elements = append(elements,"thorium");
    elements = append(elements,"protactinium");
    elements = append(elements,"uranium");
    elements = append(elements,"neptunium");
    elements = append(elements,"plutonium");
    elements = append(elements,"americium");
    elements = append(elements,"curium");
    elements = append(elements,"berkelium");
    elements = append(elements,"californium");
    elements = append(elements,"einsteinium");
    elements = append(elements,"fermium");
    elements = append(elements,"mendelevium");
    elements = append(elements,"nobelium");
    elements = append(elements,"lawrencium");
    elements = append(elements,"rutherfordium");
    elements = append(elements,"dubnium");
    elements = append(elements,"seaborgium");
    elements = append(elements,"bohrium");
    elements = append(elements,"hassium");
    elements = append(elements,"meitnerium");
    elements = append(elements,"darmstadtium");
    elements = append(elements,"roentgenium");
    elements = append(elements,"copernicium");

    minDist := 999999
    minDistIndex := -1
    for i:=0; i<len(elements); i++ {
        dist := levenshtein.Distance(elements[i], "tillium")
        if(dist < minDist) {
            minDist = dist
            minDistIndex = i
        }
    }

    fmt.Println(elements[minDistIndex])

}
