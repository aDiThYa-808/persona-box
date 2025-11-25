<script lang="ts">
    import {PUBLIC_AWS_INVOKE_URL} from "$env/static/public"

    type ChatMessage = {
        role: "user" | "assistant"
        text: string
    }

    const apiBaseUrl = PUBLIC_AWS_INVOKE_URL

    let prompt = ""
    let loading = false
    let messages : ChatMessage[] = []

    async function sendPrompt(prompt:string){
        messages = [...messages,{role:"user",text:prompt}]

        loading = true
        try{
            const response = await fetch(`${apiBaseUrl}/chat`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "message": prompt
            })
        })

        if(!response.ok){
            throw new Error(`HTTP error. Status: ${response.status}`)
        }

        const data = await response.json()
        
        messages = [...messages,{role:"assistant",text:data.response}]

    }
    catch(err){
        messages = [...messages,{role:"assistant",text:(err as Error).message}]
    }
    finally{
        loading = false
    }
}

    function handleSend(){
        if (!prompt.trim() || loading) return
        sendPrompt(prompt)
        prompt = ""
    }

</script>

<div class="flex flex-col h-screen text-gray-100">
  <div class="flex-1 overflow-y-auto px-6 py-8 space-y-4">
    {#each messages as msg, i (i)}
      <div class="flex {msg.role === 'user' ? 'justify-end' : 'justify-start'}">
        <div
          class={`max-w-[75%] px-4 py-2 rounded-2xl text-sm leading-relaxed ${
            msg.role === 'user'
              ? 'bg-indigo-600 text-white rounded-br-none shadow-md'
              : 'bg-[#1E1E20] text-gray-200 border border-gray-700 rounded-bl-none'
          }`}
        >
          {msg.text}
        </div>
      </div>
    {/each}

    {#if loading}
      <div class="flex justify-start">
        <div class="px-4 py-2 rounded-2xl text-sm bg-[#1E1E20] border border-gray-700 text-gray-400 animate-pulse">
          Thinking...
        </div>
      </div>
    {/if}
  </div>

  <div class="border-t border-gray-800 bg-[#121213] p-4 flex items-center space-x-2">
    <textarea
      bind:value={prompt}
      rows="1"
      placeholder="Type a message..."
      class="flex-1 resize-none rounded-xl bg-[#1A1A1C] border border-gray-700 px-4 py-2 text-sm text-gray-100 placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
      on:keydown={(e) => e.key === 'Enter' && !e.shiftKey && (e.preventDefault(), handleSend())}
    ></textarea>

    <button
      on:click={handleSend}
      class="bg-indigo-600 hover:bg-indigo-700 text-white rounded-xl px-4 py-2 text-sm font-medium disabled:opacity-50 transition"
      disabled={loading}
    >
      {loading ? "..." : "Send"}
    </button>
  </div>
</div>

