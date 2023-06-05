# Integrated Health Profile - IHP

A decentralised solution for digitalising medical records with an added encryption layer.

## Tech Stack used

- Golang - Backend API (for interoperability)
- Solidity - Blockchain (Ethereum contracts, hosted on Sepholia Testnet)
- InterPlanetary File System (IPFS) - Database (Best way to store big data like files)

## Problem statement

Handling medical papers like blood reports and xray records can be a tedious task. Also an average Indian spends 50 mins in a queue at a hospital (Source: NHA Reports). Implementing a centralised storage exposes a huge risk of data breach and fraudulent records.

## Solution Implemented

We are digitalising all the medical records and storing them on an Integrated Heath Profile of a user. This profile will have an IHP ID and an IHP Card with a QR Code that can be used to scan all the details and upload records. These records are stored on a IPFS Nodes with the CID being encrypted using AES and then pushed onto a decentralised chain with added encryption. This will heavily reduce the cost and also speed up the performance. As per NHA Reports, the 50 min time waiting time can now be reduced to 5 mins.
