# Posterior Network — A Minimal Bayesian Reputation Simulator

This project implements the core mechanism behind reputation systems based on Bayesian updating.  
It models a world with two types of sellers, two types of signals, and a buyer who updates their beliefs after each interaction.

## Overview

The system follows a simple structure:

- **Seller Types**
  - `Loyal`: Usually provides high-quality outcomes.
  - `Strategic`: Acts based on long-term payoff but may provide lower quality in the short run.

- **Signals**
  - `Good`: Positive outcome delivered by the seller.
  - `Bad`: Negative outcome delivered by the seller.

- **World Model**
  The world model stores likelihood values:

P(signal | seller_type)

These likelihoods represent how each seller type tends to behave.

- **Belief Updating**
The buyer maintains a prior belief:

P(Loyal), P(Strategic)

After observing a signal, the buyer updates these beliefs using Bayes’ rule:

P(type | signal) ∝ P(signal | type) * P(type)


## Files

- `computation.go` — Contains the Bayesian model, likelihoods, and belief update logic.
- `main.go` — Demonstrates how to create a model and update beliefs after a signal.

## Running the Example

1. Build the project:
 ```bash
 go build
  ```
2. Run the executable:
    ./posterior-network
You should see output similar to:
    map[0:0.6923076923 1:0.3076923077]


