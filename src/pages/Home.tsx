import { BaseUrl } from "../utils/api_config";
import axios, { AxiosRequestConfig, AxiosResponse } from "axios";
import { useEffect, useState } from "react";

type textResponse = {
  name: string
}

export function Home() {
  const [text, setText] = useState("");

  useEffect(() => {
    fetch();
  }, [])

  function fetch() {
    const options: AxiosRequestConfig = {
      url: BaseUrl + "/",
      method: "GET"
    }
    axios(options).then((res: AxiosResponse<textResponse>) => {
      const { data } = res;
      setText(data.name)
    })
  }

  return <h2>{text}</h2>;
}
