import type { RequestEvent } from "@sveltejs/kit";
import { error, isHttpError} from "@sveltejs/kit";
import { PUBLIC_AWS_INVOKE_URL } from "$env/static/public";
import { PUBLIC_ENVIRONMENT } from "$env/static/public";

export async function POST(event: RequestEvent) {
    try{
        const prodEnv = PUBLIC_ENVIRONMENT
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

        if(!data.access_token){
            throw error(res.status, 'Internal Server Error')
        }

        let secure = true
        if(prodEnv.toLowerCase() == 'local'){
            secure = false
        }

        event.cookies.set(
            'pb_access_token',
            data.access_token,
            {
                httpOnly:true,
                secure: secure,
                path:'/',
                maxAge:7*24*60*60, // 7 days
                sameSite:"lax"
            }
        )

        return new Response(JSON.stringify({success:true}),{status:200})

    }catch(err){
        if(isHttpError(err)){
            throw err
        }

        console.log(err)
        throw error(500,"Internal server error")
    }    
}