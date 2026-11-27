local value = redis.call("GET", KEYS[1])
if not value then
    return 0
end

local setting = cjson.decode(value)
local countdown = setting.count_down
if type(countdown) ~= "table" then
    return 0
end

local mediaCount = countdown.media_count
if type(mediaCount) ~= "number" or mediaCount <= 0 then
    return 0
end

countdown.media_count = math.max(0, mediaCount - tonumber(ARGV[1]))
redis.call("SET", KEYS[1], cjson.encode(setting), "KEEPTTL")
return 1
