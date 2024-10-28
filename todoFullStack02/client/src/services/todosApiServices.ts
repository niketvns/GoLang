import {apiBaseUrl} from "../utils/helperVariables";

export const getAllTodos = async () => {
    try {
        const res = await fetch(apiBaseUrl + "/todos");
        const data = await res.json();

        if(!res.ok) {
            throw new Error(data.error || "Something went wrong");
        }

        return data?.todos?.reverse() ?? [];
    } catch(error) {
        console.log(error);
    }
}