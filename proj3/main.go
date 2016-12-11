package main

import(
  "sort"
  "os"
  "fmt"
  "log"
  "bufio"
  "strings"
  conv "strconv"
  "math/rand"
  "math"
  "time"
  "github.com/davecgh/go-spew/spew"
)

var dump = spew.Dump
var print = fmt.Println
var generations = 1000*4
var actions = []int{0,1, -1} // Buy, sell, do nothing
var parentCount = 100
var offspringCount = parentCount * 2
var startingWorth = 10000.0
var lineCount = 50
var offspringSwapProportion = 0.1
var offspringMutationProportion = 0.05
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")
var IDLength = 3
var maxIndividualWorth = math.SmallestNonzeroFloat64
var minIndividualWorth = math.MaxFloat64
var wealthiestIndi = 0.0

type Record struct {
  date string
  open float64
  high float64
  low float64
  close float64
  volume int
}

type Individual struct{
  id string
  genes []int
  cash float64
  stocks int
  worth float64
}

func main(){
  rand.Seed(time.Now().UTC().UnixNano())
  records := readFile()
  reverse(records)
  // print(records)
  individuals := setup(parentCount, len(records), records)
  details(individuals, 0)
  evolve(individuals, records)
}

// Sets up initial set up parents and returns them.
func setup(individualCount int, n int, records []Record ) []Individual {
  parents := []Individual{}
  for i := 0; i < individualCount; i++{
    genes := []int{}
    for j := 0; j < n; j++{
      genes = append(genes, randomAction())
    }
    individual := Individual{id: RandStringRunes(IDLength),genes:genes, cash:startingWorth, stocks:0, worth:startingWorth}
    calculate(&individual, records)
    parents = append(parents, individual)
    genes = []int{}
  }
  return parents
}

func randomAction() int{
  return actions[rand.Int() % len(actions)]
}

// Selects survivors from a slice of Individual by means of Partial crossover (PMX)
func selectSurvivors(indis []Individual) []Individual{
  uSurivors := make(map[float64]bool)
  survivors := []Individual{}
  totalWorth := genTotalWorth(indis)
  for len(survivors) < offspringCount{
    for i := range indis{
      if rand.Float64() < (indis[i].worth / totalWorth) {
        survivors = append(survivors, indis[i])
        uSurivors[indis[i].worth] = true
        if len(survivors) >= offspringCount{
          break
        }
      }
    }
  }
  return survivors
}

// Takes a slice of Individuals and returns the summation of their worth
func genTotalWorth(indis []Individual) float64{
  totalGenWorth := 0.0
  for i := range indis{
    totalGenWorth += indis[i].worth
  }
  return totalGenWorth
}

// Sum up the cash for a slice of Individuals
func genCash(indis []Individual) float64{
  totalGenCash := 0.0
  for i := range indis{
    totalGenCash += indis[i].cash
  }
  return totalGenCash / float64(len(indis))
}

// evolve takes an original set of parents and creates offspring, and evolves them for the required number of generations
func evolve(indis []Individual, records []Record){
  for e := 1; e < generations; e++{
    sort.Sort(ByWorth(indis))
    // indis = selectSurvivors(indis) // Selected by PMX
    indis = getOffspring(indis, records)
    for _, e := range indis{
      if e.worth > maxIndividualWorth{
        maxIndividualWorth = e.worth
      }
      if e.worth > 0.0 && e.worth < minIndividualWorth{
        minIndividualWorth = e.worth
      }
    }
    details(indis, e)
    maxIndividualWorth = math.SmallestNonzeroFloat64
    minIndividualWorth = math.MaxFloat64
  }
  print("Wealthiest:", wealthiestIndi)
}

// details prints basic information for a generation of Individuals
func details(individuals []Individual, n int){
  genWorth := genTotalWorth(individuals) / float64(len(individuals))
  if len(os.Args) < 2{
    print("Generation", n, "\tavg cash:", genCash(individuals), "\tavg worth:", genWorth, "\tsize:", len(individuals), "max:", maxIndividualWorth, "min", minIndividualWorth)
  }else if os.Args[1] == "1"{
    fmt.Printf("%v,%v,%v,%v\n", n, genWorth, minIndividualWorth, maxIndividualWorth)
  }
}

func getOffspring(parents []Individual, records []Record) []Individual{
  uOffspring := make(map[float64]bool)
  offspring := []Individual{}
  for len(offspring) < offspringCount{ // While we need more offspring
    shuffle(parents) // Random mating

    for i := 1; i < len(parents); i += 2{
      o1 := parents[i]
      o2 := parents[i-1]

      //Reset o1 and o2
      o1.id, o1.cash, o1.stocks, o1.worth = RandStringRunes(IDLength), startingWorth, 0, startingWorth
      o2.id, o2.cash, o2.stocks, o2.worth = RandStringRunes(IDLength), startingWorth, 0, startingWorth

      // Iterate through genes and apply recombination techniques
      for j := 0; j < len(o1.genes); j++{
        rng := rand.Float64()
        if rng < offspringSwapProportion{
          // print("swapped")
          temp := o1.genes[j:]
          copy(o1.genes[j:], o2.genes[j:])
          copy(o2.genes[j:], temp)
        }

        // Offspring one per gene mutation
        rng = rand.Float64()
        if rng < offspringMutationProportion{
          o1.genes[j] = randomAction()
        }

        // Offspring two per gene mutation
        rng = rand.Float64()
        if rng < offspringMutationProportion{
          o2.genes[j] = randomAction()
        }
      }

      calculate(&o1, records)
      calculate(&o2, records)

      offspring = append(offspring, o1, o2)
      uOffspring[o1.worth] = true
      uOffspring[o2.worth] = true
      if len(offspring) >= offspringCount{
        break // we have enough offspring, stahp
      }
    }
  }
  return offspring
}


// calculate applies financial records to an individual's genes and sets variables on the individual.
func calculate(indi* Individual, records []Record) *Individual{
  for i, _ := range records{
    // Only sell if you have stocks, only buy if you have money
    if (indi.genes[i] < 0 && indi.stocks > 0) || (indi.genes[i] > 0){
      indi.stocks += indi.genes[i]
      indi.cash += -float64(indi.genes[i]) * records[i].close
    }
    if indi.cash < 0{
      indi.worth = 0
      return indi
    }
  }
  indi.worth = indi.cash + (float64(indi.stocks) * records[len(records) - 1].close)
  return indi
}

// Functions describing descending sorting on a slice of Individuals.
type ByWorth []Individual
func (a ByWorth) Len() int { return len(a)}
func (a ByWorth) Swap(i, j int) { a[i], a[j] = a[j], a[i]}
func (a ByWorth) Less(i, j int) bool { return a[i].worth > a[j].worth }

// reverse will reverse a slice in place.
func reverse(s []Record){
  for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
      s[i], s[j] = s[j], s[i]
  }
}

// shuffle will shuffle and array in place
func shuffle(slice []Individual){
  for i := range slice {
    j := rand.Intn(i + 1)
    slice[i], slice[j] = slice[j], slice[i]
  }
}

// RandStringRunes creates a string from letterRunes with length n.
func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

// readFile reads in data from googl.csv and returns a slice of corresponding records.
func readFile() []Record{
  records := []Record{}

  file, err := os.Open("googl.csv")
  if err != nil{
    log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)

  lineCounter := 0
  for scanner.Scan(){
    if lineCounter > 0 && len(records) < lineCount{
      strs := strings.Split(scanner.Text(), ",")
      open,_ := conv.ParseFloat(strs[1], 64)
      high,_ := conv.ParseFloat(strs[2], 64)
      low,_ :=conv.ParseFloat(strs[3], 64)
      close, _ := conv.ParseFloat(strs[4], 64)
      volume, _ := conv.Atoi(strs[5])
      record := Record{strs[0], open, high, low, close, volume}
      records = append(records, record)
    }
    lineCounter += 1
  }
  return records
}
