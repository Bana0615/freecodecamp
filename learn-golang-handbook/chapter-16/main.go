package main

import (
	"fmt"
	"math/rand"
)

// Gene: A single character (A, B, C, or D).
type Gene rune

// Chromosome: A sequence of 5 genes.
type Chromosome [5]Gene

// Individual: A chromosome and its fitness.
type Individual struct {
	Chromosome Chromosome
	Fitness    int
}

// randomGene generates a random gene.
func randomGene() Gene {
	genes := []Gene{'A', 'B', 'C', 'D'}
	return genes[rand.Intn(len(genes))]
}

// randomChromosome creates a random chromosome.
func randomChromosome() Chromosome {
	var chromosome Chromosome
	for i := 0; i < len(chromosome); i++ {
		chromosome[i] = randomGene()
	}
	return chromosome
}

// calculateFitness:  Count how many genes match the target.
func calculateFitness(chromosome Chromosome, target Chromosome) int {
	fitness := 0
	for i := 0; i < len(chromosome); i++ {
		if chromosome[i] == target[i] {
			fitness++
		}
	}
	return fitness
}

// crossover:  Simple single-point crossover (midpoint).
func crossover(parent1, parent2 Chromosome) (Chromosome, Chromosome) {
	midpoint := len(parent1) / 2
	var child1, child2 Chromosome

	// First half from parent1, second half from parent2.
	for i := 0; i < midpoint; i++ {
		child1[i] = parent1[i]
		child2[i] = parent2[i]
	}
	for i := midpoint; i < len(parent1); i++ {
		child1[i] = parent2[i]
		child2[i] = parent1[i]
	}
	return child1, child2
}

// mutate:  Randomly change a single gene with a small probability.
func mutate(chromosome Chromosome, mutationRate float64) Chromosome {
	for i := 0; i < len(chromosome); i++ {
		if rand.Float64() < mutationRate {
			chromosome[i] = randomGene() // Change to a random gene.
		}
	}
	return chromosome
}

func main() {
	target := Chromosome{'A', 'B', 'C', 'D', 'A'}
	populationSize := 10
	mutationRate := 0.1
	generations := 100

	// 1. Create Initial Population.
	population := make([]Individual, populationSize)
	for i := 0; i < populationSize; i++ {
		population[i] = Individual{
			Chromosome: randomChromosome(),
			Fitness:    0, // Fitness will be calculated later.
		}
	}

	// 2. Evolution Loop.
	for generation := 0; generation < generations; generation++ {

		// Calculate fitness for each individual.
		for i := 0; i < populationSize; i++ {
			population[i].Fitness = calculateFitness(population[i].Chromosome, target)
		}

		// Find the best individual in this generation.
		bestIndividual := population[0]
		for i := 1; i < populationSize; i++ {
			if population[i].Fitness > bestIndividual.Fitness {
				bestIndividual = population[i]
			}
		}
		fmt.Printf("Generation %d: Best: %s (Fitness: %d)\n", generation, string(bestIndividual.Chromosome[:]), bestIndividual.Fitness)

		if bestIndividual.Fitness == 5 {
			fmt.Println("Found perfect solution!")
			break
		}

		// 3. Create Next Generation (simplified selection, crossover, mutation).
		newPopulation := make([]Individual, populationSize)
		// Simple Elitism, keep top 2
		newPopulation[0] = bestIndividual
		newPopulation[1] = population[1] // Assume population[1] is second best

		for i := 2; i < populationSize; i += 2 {
			//Simple selection.  Could replace with a more sophisticated method
			parent1 := population[rand.Intn(populationSize)]
			parent2 := population[rand.Intn(populationSize)]

			// Crossover.
			child1Chromosome, child2Chromosome := crossover(parent1.Chromosome, parent2.Chromosome)

			// Mutation.
			child1Chromosome = mutate(child1Chromosome, mutationRate)
			child2Chromosome = mutate(child2Chromosome, mutationRate)

			newPopulation[i] = Individual{Chromosome: child1Chromosome}
			if i+1 < populationSize { // Avoid out-of-bounds access
				newPopulation[i+1] = Individual{Chromosome: child2Chromosome}
			}
		}

		// Replace the old population with the new one.
		population = newPopulation
	}
}
