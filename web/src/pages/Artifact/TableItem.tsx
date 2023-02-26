/**
 * The MIT License (MIT)
 *
 * Copyright © 2023 Tosone <i@tosone.cn>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

import { useRef, useState } from "react";
import { useClickAway } from 'react-use';
import { useNavigate } from 'react-router-dom';
import relativeTime from 'dayjs/plugin/relativeTime';
import dayjs from 'dayjs';

dayjs.extend(relativeTime);

export default function TableItem({ id, namespace, repository, digest, size, tags, tag_count, created_at, updated_at }: { id: number, namespace: string | undefined, repository: string | null, digest: string, size: number, tags: string[], tag_count: number, created_at: string, updated_at: string }) {
  const navigate = useNavigate();
  let [show, setShow] = useState(false);

  const ref = useRef<HTMLDivElement>() as React.MutableRefObject<HTMLDivElement>;;
  useClickAway(ref, () => {
    if (show) {
      setShow(!show);
    }
  });

  return (
    <tr className="cursor-pointer" onClick={() => {
      navigate(`/namespace/${namespace}/tag?repository=${repository}`);
    }}>
      <td className="px-6 py-4 max-w-0 w-full whitespace-nowrap text-sm font-medium text-gray-900">
        <div className="flex items-center space-x-3 lg:pl-2">
          <div className="cursor-pointer truncate hover:text-gray-600">
            {digest}
          </div>
        </div>
      </td>
      <td className="hidden md:table-cell px-4 py-4 whitespace-nowrap text-sm text-gray-500 text-right">
        {
          tags.map((tag, index) => {
            if (index < 3) {
              return (
                <span key={index} className="inline-flex items-center mx-0.5 px-1.5 py-0.5 rounded-md text-xs font-medium bg-gray-100 text-gray-800">
                  {tag.length > 10 ? tag.slice(0, 10) + '...' : tag}
                </span>
              )
            }
          })
        }
        {
          tag_count >= 3 ? (
            <span className="inline-flex items-center px-1.5 py-0.5 mx-0.5 rounded-md text-xs font-medium bg-gray-100 text-gray-800">
              ...
            </span>
          ) : (
            <></>
          )
        }
      </td>
      <td className="hidden md:table-cell px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-right">
        {tag_count}
      </td>
      <td className="hidden md:table-cell px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-right">
        {size}
      </td>
      <td className="hidden md:table-cell px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-right">
        {dayjs().to(dayjs(created_at))}
      </td>
      <td className="hidden md:table-cell px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-right">
        {dayjs().to(dayjs(updated_at))}
      </td>
      <td className="pr-3 whitespace-nowrap">
        <button
          type="button"
          className=" w-1/2  rounded-md border border-transparent bg-white font-medium text-indigo-600 hover:text-indigo-500  mr-5"
        >
          Remove
        </button>
      </td>
    </tr>
  );
}
