import { useState, type FormEvent } from 'react';
import { createShortUrl } from '../api/urlShorten';

const URLShortener = () => {
  const [url, setUrl] = useState('');
  const [expireTime, setExpireTime] = useState(24); // 默认24小时
  const [shortUrl, setShortUrl] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    
    if (!url) {
      setError('请输入有效的URL');
      return;
    }

    try {
      setLoading(true);
      setError(null);
      
      const result = await createShortUrl({
        origin_url: url,
        expire_time: expireTime
      });
      
      setShortUrl(result.short_url);
    } catch (err) {
      setError('生成短链接失败，请稍后重试');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const copyToClipboard = () => {
    if (shortUrl) {
      navigator.clipboard.writeText(shortUrl)
        .then(() => alert('链接已复制到剪贴板'))
        .catch(err => console.error('复制失败:', err));
    }
  };

  return (
    <div className="url-shortener">
      <h2>URL短链接生成器</h2>
      
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="url">输入长URL:</label>
          <input
            type="url"
            id="url"
            value={url}
            onChange={(e) => setUrl(e.target.value)}
            placeholder="https://example.com/very/long/url"
            required
          />
        </div>
        
        <div className="form-group">
          <label htmlFor="expireTime">过期时间 (小时):</label>
          <input
            type="number"
            id="expireTime"
            value={expireTime}
            onChange={(e) => setExpireTime(Number(e.target.value))}
            min="1"
            required
          />
        </div>
        
        <button type="submit" disabled={loading}>
          {loading ? '生成中...' : '生成短链接'}
        </button>
      </form>
      
      {error && <div className="error">{error}</div>}
      
      {shortUrl && (
        <div className="result">
          <h3>生成的短链接:</h3>
          <div className="short-url-container">
            <a href={shortUrl} target="_blank" rel="noopener noreferrer">
              {shortUrl}
            </a>
            <button onClick={copyToClipboard}>复制</button>
          </div>
        </div>
      )}
    </div>
  );
};

export default URLShortener;
