package ts3

type channelInfo struct {
	cid                                  int    `ms:"cid"`
	cpid                                 int    `ms:"cpid"`
	channelName                          string `ms:"channel_name"`
	channelTopic                         string `ms:"channel_topic"`
	channelDescription                   string `ms:"channel_description"`
	channelPassword                      string `ms:"channel_password"`
	channelFlagPassword                  int    `ms:"channel_flag_password"`
	channelCodecQuality                  int    `ms:"channel_codec_quality"`
	channelMaxclients                    int    `ms:"channel_maxclients"`
	channelMaxfamilyclients              int    `ms:"channel_maxfamilyclients"`
	channelOrder                         int    `ms:"channel_order"`
	channelFlagPermanent                 int    `ms:"channel_flag_permanent"`
	channelFlagSemiPermanent             int    `ms:"channel_flag_semi_permanent"`
	channelFlagTemporary                 int    `ms:"channel_flag_temporary"`
	channelFlagDefault                   int    `ms:"channel_flag_default"`
	channelFlagMaxclientsUnlimited       int    `ms:"channel_flag_maxclients_unlimited"`
	channelFlagMaxfamilyclientsUnlimited int    `ms:"channel_flag_maxfamilyclients_unlimited"`
	channelFlagMaxfamilyclientsInherited int    `ms:"channel_flag_maxfamilyclients_inherited"`
	channelNeededTalkPower               int    `ms:"channel_needed_talk_power"`
	channelNamePhonetic                  string `ms:"channel_name_phonetic"`
	channelFilepath                      string `ms:"channel_filepath"`
	channelForcedSilence                 int    `ms:"channel_forced_silence"`
	channelIconID                        int    `ms:"channel_icon_id"`
	channelCodecIsUnencrypted            int    `ms:"channel_codec_is_unencrypted"`
}

// Whoami returns information about the current connection including the currently selected virtual server.
func (c *Client) channelInfo(cid int) (*channelInfo, error) {
	i := &channelInfo{}
	if _, err := c.ExecCmd(NewCmd("channelinfo cid=4").WithResponse(&i)); err != nil {
		return nil, err
	}

	return i, nil
}
