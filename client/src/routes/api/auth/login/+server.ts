import type { RequestEvent } from "@sveltejs/kit";
import { error, isHttpError } from "@sveltejs/kit";
import { PUBLIC_AWS_INVOKE_URL } from "$env/static/public";

export async function POST(event: RequestEvent) {
    try{
        const invokeURL = PUBLIC_AWS_INVOKE_URL 

        const body = await event.request.json()
        const idToken = body.id_token

        if (!idToken){
            throw error(400,"Missing id token")
        }

        const res = await fetch(`${invokeURL}/auth/google`,{
            method: 'POST',
            headers:{
                "Content-Type":"application/json"
            },
            body: JSON.stringify({
                "id_token":idToken
            })
        })

        const data = await res.json()

        if(!res.ok){
            throw error(res.status,"Unauthorized")
        }

        if(data.error){
            console.log(data.error)
            throw error(res.status,'Unauthorized')
        }

        if(!data.sub || !data.access_token || !data.email){
            throw error(res.status, 'Internal Server Error')
        }

        return new Response(JSON.stringify(data),{status:200})

    }catch(err:unknown){
        if(isHttpError(err)){
            throw err
        }

        console.log(err)
        throw error(500,"Internal server error")
    }    
}