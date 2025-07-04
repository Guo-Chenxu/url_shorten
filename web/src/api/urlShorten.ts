interface CreateShortUrlReq {
  origin_url: string;
  expire_time: number; // 过期时间(小时)
}

interface CreateShortUrlResp {
  origin_url: string;
  short_url: string;
  expire_time: string;
}

export async function createShortUrl(req: CreateShortUrlReq): Promise<CreateShortUrlResp> {
  try {
    const response = await fetch('/api/create', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(req),
    });

    if (!response.ok) {
      throw new Error(`请求失败: ${response.status}`);
    }

    return await response.json();
  } catch (error) {
    console.error('创建短链接失败:', error);
    throw error;
  }
}
