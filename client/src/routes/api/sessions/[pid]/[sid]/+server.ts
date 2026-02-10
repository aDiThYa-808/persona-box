import { PUBLIC_AWS_INVOKE_URL } from "$env/static/public";
import { error, isHttpError, type RequestEvent } from "@sveltejs/kit";

export async function DELETE(event: RequestEvent){
    try{
        const invokeUrl = PUBLIC_AWS_INVOKE_URL
        const cookies = event.request.headers.get("Cookie")
        let headers: Record<string,string> = {}
        if(cookies){
            headers["Cookie"] = cookies
        }

        const personaid = event.params.pid
        const sessionid = event.params.sid

        const res = await fetch(`${invokeUrl}/sessions/${personaid}/${sessionid}`,{
            method : 'DELETE',
            headers: headers 
        })

        const data = await res.json()

        if(!res.ok){
            throw error(res.status,data.error)
        }

        return new Response(JSON.stringify(data),{status:res.status})

    }catch(err){
        if(isHttpError(err)){
            throw err
        }
        console.log(err)
        throw error(500, "internal server error")
    }
}