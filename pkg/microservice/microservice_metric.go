package microservice

import "miniapi/pkg/middlewares"

func (ms *Microservice) getHttpMetricsCb() middlewares.MiddlewareMetricsCb {
	return func(err error) {
		if err != nil {
			//a.metrics.ErrorHttpRequests.Inc()
		} else {
			//a.metrics.SuccessHttpRequests.Inc()
		}
	}
}
