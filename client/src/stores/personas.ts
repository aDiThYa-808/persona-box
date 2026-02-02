import type { Chat } from "$lib/types/chat";
import type { PersonaList } from "$lib/types/persona";
import { writable } from "svelte/store";

export const personas = writable<PersonaList[]>([])
export const chats = writable<Chat[]>([])