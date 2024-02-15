# Golang Vite Integration Example

### About

This is a simplified version of my personal project structure.

This repository contains everything needed to get started on a project with **Go** using **Vite**, **HTMX**, **Tailwind**, and of course **Templ**.

### How to use

You need `air`, `pnpm`, `go` and `make` to use this project.

Don't forget to fetch Node depencies with `pnpm`

To start `vite` and `air` at the same time, use `make dev`

### In the future

In the future, I will create a tool called ***GoMaker*** that will be able to generate a similar structure to this repository with a command, 
but also integrate the database module, scheduler, mail *(with reactEmail designed to be templated with the golang module)*, and unit testing among others.

The final idea with this CLI is to be able to easily start a project with only the basics and then copy my ready-made modules so that the developer can just modify them as they wish at the end. 
(This is totally inspired by shadcn/ui even though I've never touched it).
