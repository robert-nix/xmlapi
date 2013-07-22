# xmlapi

Made to solve the problem of getting start and end times relative to a Twitch
broadcast's recorded_at times for a Twitch highlight VoD in a format consumable
by JavaScript.

Sample config for that purpose (xmlapi.json):

    {
      "http_port": 9003,
      "remote_base_url": "http://api.justin.tv/api",
      "remote_suffix": ".xml",
      "endpoints": {
        "/broadcast/by_chapter": ["archives.bracket_start","archives.bracket_end"]
      }
    }
