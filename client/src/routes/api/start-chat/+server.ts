import { PUBLIC_AWS_INVOKE_URL } from "$env/static/public";
import { error, isHttpError, type RequestEvent } from "@sveltejs/kit";

export async function POST(events: RequestEvent){
    try{
        const cookies = events.request.headers.get('Cookie')
        const rbody = events.request.json()
        const invokeUrl = PUBLIC_AWS_INVOKE_URL
        const headers: Record<string,string> = {
            'Content-Type': 'application/json'
        }
        if(cookies){
            headers['Cookie'] = cookies
        }

        const res = await fetch(`${invokeUrl}/start-chat`,{
            method:'POST',
            headers:headers,
            body: JSON.stringify(rbody)
        })

        if(!res.ok){
            throw error(res.status,res.statusText)
        }

        const chat = await res.json()
        
        return new Response(JSON.stringify(chat),{status:200})

    }catch(err){
        if(isHttpError(err)){
            throw err
        }

        console.log(err)
        throw error(500,"internal server error")
    }
}