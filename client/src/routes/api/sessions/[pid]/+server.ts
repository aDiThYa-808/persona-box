import { error, isHttpError, type RequestEvent } from "@sveltejs/kit";
import { PUBLIC_AWS_INVOKE_URL } from "$env/static/public";
import type Chat from "$lib/components/chat.svelte";

export async function GET(events: RequestEvent){
    try{
        const invokeUrl = PUBLIC_AWS_INVOKE_URL
        const cookies = events.request.headers.get("Cookie")
        const personaid = events.params.pid
        let headers: Record<string,string> = {}
        if(cookies){
            headers['Cookie'] = cookies
        }

        const res = await fetch(`${invokeUrl}/sessions/${personaid}`,{
            method: 'GET',
            headers: headers
        })

        if(!res.ok){
            let data = await res.json()
            throw error(res.status,data.error)
        }

        const data: Chat[] = await res.json()

        return new Response(JSON.stringify(data),{status:200})

    }catch(err){
        if(isHttpError(err)){
            throw err
        }

        console.log(err)
        throw error(500,"internal server error")
    }
}