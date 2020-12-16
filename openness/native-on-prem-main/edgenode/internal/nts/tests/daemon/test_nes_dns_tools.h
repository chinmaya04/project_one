/*
 * SPDX-License-Identifier: Apache-2.0
 * Copyright (c) 2019 Intel Corporation
 */

#ifndef TEST_NES_DNS_TOOLS_H
#define	TEST_NES_DNS_TOOLS_H

#include <CUnit/CUnit.h>

int init_suite_nes_dns_tools(void);
int cleanup_suite_nes_dns_tools(void);

void add_nes_dns_tools_suite_to_registry(void);

#endif /* TEST_NES_DNS_TOOLS_H */
