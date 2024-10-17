import { BaseUrl } from "../utils/api_config";
import axios, { AxiosRequestConfig, AxiosResponse } from "axios";
import { useState, useEffect } from "react";
import { Link } from "react-router-dom";

type GachaInfo = {
  id: number,
  name: string,
  puS: string,
  puA_name_0: string,
  puA_name_1: string,
  startDate: string,
  endDate: string
}

export function Home() {
  const [gachaList, setGachaList] = useState<GachaInfo[]>([]);

  useEffect(() => {
    fetchList();
  }, [])

  function fetchList() {
    const options: AxiosRequestConfig = {
      baseURL: BaseUrl,
      url: `/gacha`,
      method: "get"
    }

    axios(options).then((res: AxiosResponse<{results: GachaInfo[]}>) => {
      const { data } = res;
      setGachaList(data.results);
    })
  }


  return <table>
    <tbody>
    {
      gachaList.map((i: GachaInfo) => {
        return <tr>
          <td><Link to={`/gacha/${i.id}`}>{i.name}</Link></td>
          <td>{i.puS}</td>
          <td>{i.puA_name_0}</td>
          <td>{i.puA_name_1}</td>
          <td>{`${i.startDate} ~ ${i.endDate}`}</td>
        </tr>
      })
    }
    </tbody>
  </table>
}
