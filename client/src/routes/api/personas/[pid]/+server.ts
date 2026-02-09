import { PUBLIC_AWS_INVOKE_URL } from "$env/static/public";
import { error, isHttpError, type RequestEvent } from "@sveltejs/kit";

export async function DELETE(events: RequestEvent){
    try{
        const invokeUrl = PUBLIC_AWS_INVOKE_URL
        const cookies = events.request.headers.get("Cookie")
        const personaid = events.params.pid
        let headers: Record<string,string> = {}
        if(cookies){
            headers['Cookie'] = cookies
        }

        const res = await fetch(`${invokeUrl}/personas/${personaid}`,{
            method: 'DELETE',
            headers: headers
        })

        if(!res.ok){
            let data = await res.json()
            throw error(res.status,data.error)
        }

        const data = await res.json()

        return new Response(JSON.stringify(data),{status: res.status})
        
    }catch(err){
        if(isHttpError(err)){
            throw err
        }
        console.log(err)
        throw error(500, "internal server error")
    }
}