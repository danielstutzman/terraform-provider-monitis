package main

import (
	"fmt"
	"github.com/danielstutzman/go-monitis"
	"github.com/hashicorp/terraform/helper/schema"
	"strconv"
	"strings"
)

func resource_monitis_external_monitor() *schema.Resource {
	return &schema.Resource{
		Create: resource_monitis_external_monitor_create,
		Read:   resource_monitis_external_monitor_read,
		Update: resource_monitis_external_monitor_update,
		Delete: resource_monitis_external_monitor_delete,
		Importer: &schema.ResourceImporter{
			State: resource_monitis_external_monitor_import,
		},

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"detailed_test_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"location_ids": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tag": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"over_ssl": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"post_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"content_match_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"content_match_flag": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"params": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"uptime_sla": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"response_sla": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"basic_auth_user": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				ForceNew:  true,
			},
			"basic_auth_pass": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				ForceNew:  true,
			},
			"header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sni": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_version_1_1": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"user_agent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"order_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"is_ipv6": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"start_date": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func convertDataToAddOptions(d *schema.ResourceData) *monitis.AddExternalMonitorOptions {
	opts := monitis.AddExternalMonitorOptions{}

	if type_, ok := d.GetOk("type"); ok {
		opts.Type = monitis.String(type_.(string))
	}
	if detailed_test_type, ok := d.GetOk("detailed_test_type"); ok {
		opts.DetailedTestType = monitis.Int(detailed_test_type.(int))
	}
	if name, ok := d.GetOk("name"); ok {
		opts.Name = monitis.String(name.(string))
	}
	if url, ok := d.GetOk("url"); ok {
		opts.Url = monitis.String(url.(string))
	}
	if timeout, ok := d.GetOk("timeout"); ok {
		opts.Timeout = monitis.Int(timeout.(int))
	}
	if tag, ok := d.GetOk("tag"); ok {
		opts.Tag = monitis.String(tag.(string))
	}
	if interval, ok := d.GetOk("interval"); ok {
		opts.Interval = monitis.Int(interval.(int))
	}
	if over_ssl, ok := d.GetOk("over_ssl"); ok {
		opts.OverSsl = monitis.BoolToInt(over_ssl.(bool))
	}
	if post_data, ok := d.GetOk("post_data"); ok {
		opts.PostData = monitis.String(post_data.(string))
	}
	if content_match_string, ok := d.GetOk("content_match_string"); ok {
		opts.ContentMatchString = monitis.String(content_match_string.(string))
	}
	if content_match_flag, ok := d.GetOk("content_match_flag"); ok {
		opts.ContentMatchFlag = monitis.Int(content_match_flag.(int))
	}
	if params, ok := d.GetOk("params"); ok {
		opts.Params = monitis.String(params.(string))
	}
	if uptime_sla, ok := d.GetOk("uptime_sla"); ok {
		opts.UptimeSla = monitis.Int(uptime_sla.(int))
	}
	if response_sla, ok := d.GetOk("response_sla"); ok {
		opts.ResponseSla = monitis.Int(response_sla.(int))
	}
	if basic_auth_user, ok := d.GetOk("basic_auth_user"); ok {
		opts.BasicAuthUser = monitis.String(basic_auth_user.(string))
	}
	if basic_auth_pass, ok := d.GetOk("basic_auth_pass"); ok {
		opts.BasicAuthPass = monitis.String(basic_auth_pass.(string))
	}
	if header, ok := d.GetOk("header"); ok {
		opts.Header = monitis.String(header.(string))
	}
	if sni, ok := d.GetOk("sni"); ok {
		opts.Sni = monitis.BoolToInt(sni.(bool))
	}
	if is_version_1_1, ok := d.GetOk("is_version_1_1"); ok {
		opts.IsVersion_1_1 = monitis.BoolToInt(is_version_1_1.(bool))
	}
	if user_agent, ok := d.GetOk("user_agent"); ok {
		opts.UserAgent = monitis.String(user_agent.(string))
	}
	if order_id, ok := d.GetOk("order_id"); ok {
		opts.OrderId = monitis.Int(order_id.(int))
	}
	if is_ipv6, ok := d.GetOk("is_ipv6"); ok {
		opts.IsIpv6 = monitis.BoolToInt(is_ipv6.(bool))
	}

	location_ids := make([]string, d.Get("location_ids.#").(int))
	for i, location_id := range d.Get("location_ids").([]interface{}) {
		location_ids[i] = strconv.Itoa(location_id.(int))
	}
	opts.LocationIds = monitis.String(strings.Join(location_ids, ","))

	return &opts
}

func resource_monitis_external_monitor_create(d *schema.ResourceData, m interface{}) error {
	auth := m.(*monitis.Auth)
	opts := convertDataToAddOptions(d)

	monitorData, err := auth.AddExternalMonitor(opts)
	if err != nil {
		return fmt.Errorf("Error from AddExternalMonitor: %s", err)
	}

	d.SetId(strconv.Itoa(monitorData.TestId))

	return err
}

func resource_monitis_external_monitor_read(data *schema.ResourceData,
	meta interface{}) error {

	auth := meta.(*monitis.Auth)

	testId := data.Id()
	info, err := auth.GetExternalMonitorInfo(testId, nil)
	if err != nil {
		return fmt.Errorf("Error from GetExternalMonitorInfo: %s", err)
	}

	locationIds := []int{}
	var checkInterval int
	for _, location := range info.Locations {
		locationIds = append(locationIds, location.Id)
		if checkInterval == 0 {
			checkInterval = location.CheckInterval
		} else if checkInterval != location.CheckInterval {
			return fmt.Errorf("Inconsistent CheckIntervals: %d and %d",
				checkInterval, location.CheckInterval)
		}
	}

	params := []string{}
	var user_agent string
	var is_ipv6 bool
	var is_version_1_1 bool
	var ssl_version string
	var header string
	var sni bool
	for paramName, paramValue := range info.Params {
		if paramName == "useragent" {
			user_agent = paramValue.(string)
		} else if paramName == "isIPv6" {
			is_ipv6 = (paramValue.(string) == "1")
		} else if paramName == "isversion_1_1" {
			is_version_1_1 = (paramValue.(string) == "1")
		} else if paramName == "sslVersion" {
			ssl_version = paramValue.(string)
		} else if paramName == "header" {
			header = paramValue.(string)
		} else if paramName == "sni" {
			sni = (paramValue.(string) == "1")
		} else {
			params = append(params, fmt.Sprintf("%s=%s", paramName, paramValue))
		}
	}

	data.Set("timeout", info.Timeout)
	data.Set("start_date", info.StartDateParsed)
	data.Set("type", info.Type)
	data.Set("post_data", info.PostData)
	data.Set("content_match_string", info.MatchText)
	data.Set("content_match_flag", info.Match)
	data.Set("params", strings.Join(params, "&"))
	data.Set("tag", info.Tag)
	data.Set("detailed_test_type", info.DetailedType)
	data.Set("url", info.Url)
	data.Set("name", info.Name)
	data.Set("location_ids", locationIds)
	data.Set("interval", checkInterval)
	data.Set("user_agent", user_agent)
	data.Set("is_ipv6", is_ipv6)
	data.Set("is_version_1_1", is_version_1_1)
	data.Set("header", header)
	data.Set("sni", sni)
	_ = ssl_version
	// TODO: groups, isSuspended, isPhmon, sslVersion

	return nil
}

func convertDataToEditOptions(d *schema.ResourceData) *monitis.EditExternalMonitorOptions {
	opts := monitis.EditExternalMonitorOptions{}

	if name, ok := d.GetOk("name"); ok {
		opts.Name = monitis.String(name.(string))
	}
	if url, ok := d.GetOk("url"); ok {
		opts.Url = monitis.String(url.(string))
	}
	if timeout, ok := d.GetOk("timeout"); ok {
		opts.Timeout = monitis.Int(timeout.(int))
	}
	if tag, ok := d.GetOk("tag"); ok {
		opts.Tag = monitis.String(tag.(string))
	}
	if content_match_string, ok := d.GetOk("content_match_string"); ok {
		opts.ContentMatchString = monitis.String(content_match_string.(string))
	}
	if content_match_flag, ok := d.GetOk("content_match_flag"); ok {
		opts.ContentMatchFlag = monitis.Int(content_match_flag.(int))
	}
	if uptime_sla, ok := d.GetOk("uptime_sla"); ok {
		opts.UptimeSla = monitis.Int(uptime_sla.(int))
	}
	if response_sla, ok := d.GetOk("response_sla"); ok {
		opts.ResponseSla = monitis.Int(response_sla.(int))
	}
	if header, ok := d.GetOk("header"); ok {
		opts.Header = monitis.String(header.(string))
	}
	if sni, ok := d.GetOk("sni"); ok {
		opts.Sni = monitis.BoolToInt(sni.(bool))
	}
	if is_version_1_1, ok := d.GetOk("is_version_1_1"); ok {
		opts.IsVersion_1_1 = monitis.BoolToInt(is_version_1_1.(bool))
	}
	if user_agent, ok := d.GetOk("user_agent"); ok {
		opts.UserAgent = monitis.String(user_agent.(string))
	}
	if order_id, ok := d.GetOk("order_id"); ok {
		opts.OrderId = monitis.Int(order_id.(int))
	}
	if is_ipv6, ok := d.GetOk("is_ipv6"); ok {
		opts.IsIpv6 = monitis.BoolToInt(is_ipv6.(bool))
	}

	locationIdIntervalPairs := make([]string, d.Get("location_ids.#").(int))
	interval := d.Get("interval").(int)
	for i, location_id := range d.Get("location_ids").([]interface{}) {
		locationIdIntervalPairs[i] = fmt.Sprintf("%d-%d", location_id, interval)
	}
	opts.LocationIdIntervalPairs =
		monitis.String(strings.Join(locationIdIntervalPairs, ","))

	return &opts
}

func resource_monitis_external_monitor_update(d *schema.ResourceData, m interface{}) error {

	auth := m.(*monitis.Auth)
	testId := d.Id()
	opts := convertDataToEditOptions(d)

	err := auth.EditExternalMonitor(testId, opts)
	if err != nil {
		return fmt.Errorf("Error from EditExternalMonitor: %s", err)
	}

	return nil
}

func resource_monitis_external_monitor_delete(d *schema.ResourceData, m interface{}) error {
	auth := m.(*monitis.Auth)
	testId := d.Id()

	err := auth.DeleteExternalMonitors(testId, nil)

	return err
}

func resource_monitis_external_monitor_import(data *schema.ResourceData,
	meta interface{}) ([]*schema.ResourceData, error) {

	if err := resource_monitis_external_monitor_read(data, meta); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{data}, nil
}
