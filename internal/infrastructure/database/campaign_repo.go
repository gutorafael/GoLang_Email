package database

import "email/internal/domain/campaign"

type CampaignRepo struct {
	campaign []campaign.Campaign
}

func (c *CampaignRepo) Save(campaign *campaign.Campaign) error{
	c.campaign = append(c.campaign, *campaign)
	return nil
}

func (c *CampaignRepo) GetAll() []campaign.Campaign{
	return c.campaign
}