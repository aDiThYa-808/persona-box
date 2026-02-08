import type { Chat } from "$lib/types/chat";
import type { Message } from "$lib/types/message";
import type { PersonaList } from "$lib/types/persona";
import type { User } from "$lib/types/user";
import { writable } from "svelte/store";

export const user = writable<User>()
export const personas = writable<PersonaList[]>([])
export const chats = writable<Chat[]>([])
export const messages = writable<Message[]>([])