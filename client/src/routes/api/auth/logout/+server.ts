import { type RequestEvent } from "@sveltejs/kit";

export async function GET(event : RequestEvent){
    event.cookies.delete('pb_access_token',{path:"/"})
    return new Response(JSON.stringify({message:"deleted cookie"}),{status: 200})
}