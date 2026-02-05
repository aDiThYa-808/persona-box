import { PUBLIC_AWS_INVOKE_URL } from "$env/static/public";
import type { Message } from "$lib/types/message";
import { error, isHttpError, type RequestEvent } from "@sveltejs/kit";

export async function GET(events: RequestEvent){
    try{
        const cookies = events.request.headers.get("Cookie")
        const sessionid = events.url.searchParams.get("sessionid")
        const invokeUrl = PUBLIC_AWS_INVOKE_URL
        const headers: Record<string,string> ={
            'Content-Type':'application/json'
        }
        if(cookies){
            headers['Cookie'] = cookies
        }

        const res = await fetch(`${invokeUrl}/messages?sessionid=${sessionid}`,{
            method:'GET',
            headers:headers
        })

        if(!res.ok){
            let data = await res.json()
            throw error(res.status,data.error)
        }

        const messages: Message[] = await  res.json()

        return new Response(JSON.stringify(messages),{status:200})
    }catch(err){
        console.log(err)
        throw err
    }
}