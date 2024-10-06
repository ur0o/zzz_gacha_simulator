import { BaseUrl } from "../utils/api_config";
import axios, { AxiosRequestConfig, AxiosResponse } from "axios";
import { useState } from "react";

type GachaResult = {
  name: string,
  rank: string,
}

export function Home() {
  const [text, setText] = useState("");

  useEffect(() => {
    fetch();
  }, [])

  function fetch() {
    const options: AxiosRequestConfig = {
      baseURL: BaseUrl,
      url: `/gacha/1/draw`,
      method: "get",
      // params: gachastate
      params: {
        offsetS: 0,
        offsetA: 0,
        fixedS: false,
        fixedA: false,
        gacha_id: 1
      }
    }
    axios(options).then((res: AxiosResponse<GachaResult[]>) => {
      const { data } = res;
      setResult(data)
      setHistory(history.concat(result).slice(0, 90))
    })
  }

  return <div>
    <button onClick={fetch}>ガチャる</button>
    {
      result.length != 0 &&
        <table>
          <thead>
            <th>
              <td>レアリティ</td>
              <td>名前</td>
            </th>
          </thead>
          <tbody>
            {result.map((res, _) => {
              return <tr>
                <td>{res.rank}</td>
                <td>{res.name}</td>
              </tr>
            })}
          </tbody>
        </table>
    }
  </div>
}
