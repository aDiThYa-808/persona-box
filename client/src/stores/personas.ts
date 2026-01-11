import type { PersonaList } from "$lib/types/persona";
import { writable } from "svelte/store";

export const personas = writable<PersonaList[]>([])