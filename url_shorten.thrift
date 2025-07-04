namespace go url_shorten

struct BaseResp {
    1: i32 code
    2: string msg
}

struct CreateShortUrlReq {
    1: string origin_url
    2: i64 expire_time // 过期时间 单位：小时
}

struct CreateShortUrlResp {
    1: string origin_url
    2: string short_url
    3: string expire_time // 过期时间
}

struct QueryReq {
    1: string code (api.path="code")
}

service URLShortenHandler {
    CreateShortUrlResp CreateShortURL(1: CreateShortUrlReq req) (api.post="/api/create")
    BaseResp Query(1: QueryReq req) (api.get="/s/:code")
}