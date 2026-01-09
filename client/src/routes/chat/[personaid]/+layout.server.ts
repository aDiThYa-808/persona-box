import type { LayoutServerLoad } from "../$types";

export const load: LayoutServerLoad= ({params})=>{
    const {personaid,chatid} = params

    return {personaid,chatid}
}