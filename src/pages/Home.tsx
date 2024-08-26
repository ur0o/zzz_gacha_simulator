import { BaseUrl } from "../utils/api_config";
import axios, { AxiosRequestConfig, AxiosResponse } from "axios";
import { useEffect, useState } from "react";

type textResponse = {
  text: string
}

export function Home() {
  const [text, setText] = useState("");

  useEffect(() => {
    fetchText();
  }, [])

  function fetchText() {
    const options: AxiosRequestConfig = {
      url: BaseUrl + "/",
      method: "GET"
    }
    axios(options).then((res: AxiosResponse<textResponse>) => {
      const { data } = res;
      setText(data.text)
    })
  }

  return <h2>{text}</h2>;
}
