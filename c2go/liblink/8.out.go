package main

// Inferno utils/8l/span.c
// http://code.google.com/p/inferno-os/source/browse/utils/8l/span.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
// Instruction layout.
// Inferno utils/8c/8.out.h
// http://code.google.com/p/inferno-os/source/browse/utils/8c/8.out.h
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
const (
	AXXX_8 = iota
	AAAA_8
	AAAD_8
	AAAM_8
	AAAS_8
	AADCB_8
	AADCL_8
	AADCW_8
	AADDB_8
	AADDL_8
	AADDW_8
	AADJSP_8
	AANDB_8
	AANDL_8
	AANDW_8
	AARPL_8
	ABOUNDL_8
	ABOUNDW_8
	ABSFL_8
	ABSFW_8
	ABSRL_8
	ABSRW_8
	ABTL_8
	ABTW_8
	ABTCL_8
	ABTCW_8
	ABTRL_8
	ABTRW_8
	ABTSL_8
	ABTSW_8
	ABYTE_8
	ACALL_8
	ACLC_8
	ACLD_8
	ACLI_8
	ACLTS_8
	ACMC_8
	ACMPB_8
	ACMPL_8
	ACMPW_8
	ACMPSB_8
	ACMPSL_8
	ACMPSW_8
	ADAA_8
	ADAS_8
	ADATA_8
	ADECB_8
	ADECL_8
	ADECW_8
	ADIVB_8
	ADIVL_8
	ADIVW_8
	AENTER_8
	AGLOBL_8
	AGOK_8
	AHISTORY_8
	AHLT_8
	AIDIVB_8
	AIDIVL_8
	AIDIVW_8
	AIMULB_8
	AIMULL_8
	AIMULW_8
	AINB_8
	AINL_8
	AINW_8
	AINCB_8
	AINCL_8
	AINCW_8
	AINSB_8
	AINSL_8
	AINSW_8
	AINT_8
	AINTO_8
	AIRETL_8
	AIRETW_8
	AJCC_8
	AJCS_8
	AJCXZL_8
	AJCXZW_8
	AJEQ_8
	AJGE_8
	AJGT_8
	AJHI_8
	AJLE_8
	AJLS_8
	AJLT_8
	AJMI_8
	AJMP_8
	AJNE_8
	AJOC_8
	AJOS_8
	AJPC_8
	AJPL_8
	AJPS_8
	ALAHF_8
	ALARL_8
	ALARW_8
	ALEAL_8
	ALEAW_8
	ALEAVEL_8
	ALEAVEW_8
	ALOCK_8
	ALODSB_8
	ALODSL_8
	ALODSW_8
	ALONG_8
	ALOOP_8
	ALOOPEQ_8
	ALOOPNE_8
	ALSLL_8
	ALSLW_8
	AMOVB_8
	AMOVL_8
	AMOVW_8
	AMOVQ_8
	AMOVBLSX_8
	AMOVBLZX_8
	AMOVBWSX_8
	AMOVBWZX_8
	AMOVWLSX_8
	AMOVWLZX_8
	AMOVSB_8
	AMOVSL_8
	AMOVSW_8
	AMULB_8
	AMULL_8
	AMULW_8
	ANAME_8
	ANEGB_8
	ANEGL_8
	ANEGW_8
	ANOP_8
	ANOTB_8
	ANOTL_8
	ANOTW_8
	AORB_8
	AORL_8
	AORW_8
	AOUTB_8
	AOUTL_8
	AOUTW_8
	AOUTSB_8
	AOUTSL_8
	AOUTSW_8
	APAUSE_8
	APOPAL_8
	APOPAW_8
	APOPFL_8
	APOPFW_8
	APOPL_8
	APOPW_8
	APUSHAL_8
	APUSHAW_8
	APUSHFL_8
	APUSHFW_8
	APUSHL_8
	APUSHW_8
	ARCLB_8
	ARCLL_8
	ARCLW_8
	ARCRB_8
	ARCRL_8
	ARCRW_8
	AREP_8
	AREPN_8
	ARET_8
	AROLB_8
	AROLL_8
	AROLW_8
	ARORB_8
	ARORL_8
	ARORW_8
	ASAHF_8
	ASALB_8
	ASALL_8
	ASALW_8
	ASARB_8
	ASARL_8
	ASARW_8
	ASBBB_8
	ASBBL_8
	ASBBW_8
	ASCASB_8
	ASCASL_8
	ASCASW_8
	ASETCC_8
	ASETCS_8
	ASETEQ_8
	ASETGE_8
	ASETGT_8
	ASETHI_8
	ASETLE_8
	ASETLS_8
	ASETLT_8
	ASETMI_8
	ASETNE_8
	ASETOC_8
	ASETOS_8
	ASETPC_8
	ASETPL_8
	ASETPS_8
	ACDQ_8
	ACWD_8
	ASHLB_8
	ASHLL_8
	ASHLW_8
	ASHRB_8
	ASHRL_8
	ASHRW_8
	ASTC_8
	ASTD_8
	ASTI_8
	ASTOSB_8
	ASTOSL_8
	ASTOSW_8
	ASUBB_8
	ASUBL_8
	ASUBW_8
	ASYSCALL_8
	ATESTB_8
	ATESTL_8
	ATESTW_8
	ATEXT_8
	AVERR_8
	AVERW_8
	AWAIT_8
	AWORD_8
	AXCHGB_8
	AXCHGL_8
	AXCHGW_8
	AXLAT_8
	AXORB_8
	AXORL_8
	AXORW_8
	AFMOVB_8
	AFMOVBP_8
	AFMOVD_8
	AFMOVDP_8
	AFMOVF_8
	AFMOVFP_8
	AFMOVL_8
	AFMOVLP_8
	AFMOVV_8
	AFMOVVP_8
	AFMOVW_8
	AFMOVWP_8
	AFMOVX_8
	AFMOVXP_8
	AFCOMB_8
	AFCOMBP_8
	AFCOMD_8
	AFCOMDP_8
	AFCOMDPP_8
	AFCOMF_8
	AFCOMFP_8
	AFCOMI_8
	AFCOMIP_8
	AFCOML_8
	AFCOMLP_8
	AFCOMW_8
	AFCOMWP_8
	AFUCOM_8
	AFUCOMI_8
	AFUCOMIP_8
	AFUCOMP_8
	AFUCOMPP_8
	AFADDDP_8
	AFADDW_8
	AFADDL_8
	AFADDF_8
	AFADDD_8
	AFMULDP_8
	AFMULW_8
	AFMULL_8
	AFMULF_8
	AFMULD_8
	AFSUBDP_8
	AFSUBW_8
	AFSUBL_8
	AFSUBF_8
	AFSUBD_8
	AFSUBRDP_8
	AFSUBRW_8
	AFSUBRL_8
	AFSUBRF_8
	AFSUBRD_8
	AFDIVDP_8
	AFDIVW_8
	AFDIVL_8
	AFDIVF_8
	AFDIVD_8
	AFDIVRDP_8
	AFDIVRW_8
	AFDIVRL_8
	AFDIVRF_8
	AFDIVRD_8
	AFXCHD_8
	AFFREE_8
	AFLDCW_8
	AFLDENV_8
	AFRSTOR_8
	AFSAVE_8
	AFSTCW_8
	AFSTENV_8
	AFSTSW_8
	AF2XM1_8
	AFABS_8
	AFCHS_8
	AFCLEX_8
	AFCOS_8
	AFDECSTP_8
	AFINCSTP_8
	AFINIT_8
	AFLD1_8
	AFLDL2E_8
	AFLDL2T_8
	AFLDLG2_8
	AFLDLN2_8
	AFLDPI_8
	AFLDZ_8
	AFNOP_8
	AFPATAN_8
	AFPREM_8
	AFPREM1_8
	AFPTAN_8
	AFRNDINT_8
	AFSCALE_8
	AFSIN_8
	AFSINCOS_8
	AFSQRT_8
	AFTST_8
	AFXAM_8
	AFXTRACT_8
	AFYL2X_8
	AFYL2XP1_8
	AEND_8
	ADYNT__8
	AINIT__8
	ASIGNAME_8
	ACMPXCHGB_8
	ACMPXCHGL_8
	ACMPXCHGW_8
	ACMPXCHG8B_8
	ACPUID_8
	ARDTSC_8
	AXADDB_8
	AXADDL_8
	AXADDW_8
	ACMOVLCC_8
	ACMOVLCS_8
	ACMOVLEQ_8
	ACMOVLGE_8
	ACMOVLGT_8
	ACMOVLHI_8
	ACMOVLLE_8
	ACMOVLLS_8
	ACMOVLLT_8
	ACMOVLMI_8
	ACMOVLNE_8
	ACMOVLOC_8
	ACMOVLOS_8
	ACMOVLPC_8
	ACMOVLPL_8
	ACMOVLPS_8
	ACMOVWCC_8
	ACMOVWCS_8
	ACMOVWEQ_8
	ACMOVWGE_8
	ACMOVWGT_8
	ACMOVWHI_8
	ACMOVWLE_8
	ACMOVWLS_8
	ACMOVWLT_8
	ACMOVWMI_8
	ACMOVWNE_8
	ACMOVWOC_8
	ACMOVWOS_8
	ACMOVWPC_8
	ACMOVWPL_8
	ACMOVWPS_8
	AFCMOVCC_8
	AFCMOVCS_8
	AFCMOVEQ_8
	AFCMOVHI_8
	AFCMOVLS_8
	AFCMOVNE_8
	AFCMOVNU_8
	AFCMOVUN_8
	ALFENCE_8
	AMFENCE_8
	ASFENCE_8
	AEMMS_8
	APREFETCHT0_8
	APREFETCHT1_8
	APREFETCHT2_8
	APREFETCHNTA_8
	ABSWAPL_8
	AUNDEF_8
	AADDPD_8
	AADDPS_8
	AADDSD_8
	AADDSS_8
	AANDNPD_8
	AANDNPS_8
	AANDPD_8
	AANDPS_8
	ACMPPD_8
	ACMPPS_8
	ACMPSD_8
	ACMPSS_8
	ACOMISD_8
	ACOMISS_8
	ACVTPL2PD_8
	ACVTPL2PS_8
	ACVTPD2PL_8
	ACVTPD2PS_8
	ACVTPS2PL_8
	ACVTPS2PD_8
	ACVTSD2SL_8
	ACVTSD2SS_8
	ACVTSL2SD_8
	ACVTSL2SS_8
	ACVTSS2SD_8
	ACVTSS2SL_8
	ACVTTPD2PL_8
	ACVTTPS2PL_8
	ACVTTSD2SL_8
	ACVTTSS2SL_8
	ADIVPD_8
	ADIVPS_8
	ADIVSD_8
	ADIVSS_8
	AMASKMOVOU_8
	AMAXPD_8
	AMAXPS_8
	AMAXSD_8
	AMAXSS_8
	AMINPD_8
	AMINPS_8
	AMINSD_8
	AMINSS_8
	AMOVAPD_8
	AMOVAPS_8
	AMOVO_8
	AMOVOU_8
	AMOVHLPS_8
	AMOVHPD_8
	AMOVHPS_8
	AMOVLHPS_8
	AMOVLPD_8
	AMOVLPS_8
	AMOVMSKPD_8
	AMOVMSKPS_8
	AMOVNTO_8
	AMOVNTPD_8
	AMOVNTPS_8
	AMOVSD_8
	AMOVSS_8
	AMOVUPD_8
	AMOVUPS_8
	AMULPD_8
	AMULPS_8
	AMULSD_8
	AMULSS_8
	AORPD_8
	AORPS_8
	APADDQ_8
	APAND_8
	APCMPEQB_8
	APMAXSW_8
	APMAXUB_8
	APMINSW_8
	APMINUB_8
	APMOVMSKB_8
	APSADBW_8
	APSUBB_8
	APSUBL_8
	APSUBQ_8
	APSUBSB_8
	APSUBSW_8
	APSUBUSB_8
	APSUBUSW_8
	APSUBW_8
	APUNPCKHQDQ_8
	APUNPCKLQDQ_8
	APXOR_8
	ARCPPS_8
	ARCPSS_8
	ARSQRTPS_8
	ARSQRTSS_8
	ASQRTPD_8
	ASQRTPS_8
	ASQRTSD_8
	ASQRTSS_8
	ASUBPD_8
	ASUBPS_8
	ASUBSD_8
	ASUBSS_8
	AUCOMISD_8
	AUCOMISS_8
	AUNPCKHPD_8
	AUNPCKHPS_8
	AUNPCKLPD_8
	AUNPCKLPS_8
	AXORPD_8
	AXORPS_8
	AAESENC_8
	APINSRD_8
	APSHUFB_8
	AUSEFIELD_8
	ATYPE_8
	AFUNCDATA_8
	APCDATA_8
	ACHECKNIL_8
	AVARDEF_8
	AVARKILL_8
	ADUFFCOPY_8
	ADUFFZERO_8
	ALAST_8
)

const (
	D_AL_8 = 0 + iota
	D_CL_8
	D_DL_8
	D_BL_8
	D_AH_8 = 4 + iota - 4
	D_CH_8
	D_DH_8
	D_BH_8
	D_AX_8 = 8 + iota - 8
	D_CX_8
	D_DX_8
	D_BX_8
	D_SP_8
	D_BP_8
	D_SI_8
	D_DI_8
	D_F0_8 = 16
	D_F7_8 = D_F0_8 + 7
	D_CS_8 = 24 + iota - 18
	D_SS_8
	D_DS_8
	D_ES_8
	D_FS_8
	D_GS_8
	D_GDTR_8
	D_IDTR_8
	D_LDTR_8
	D_MSW_8
	D_TASK_8
	D_CR_8 = 35
	D_DR_8 = 43
	D_TR_8 = 51
	D_X0_8 = 59 + iota - 32
	D_X1_8
	D_X2_8
	D_X3_8
	D_X4_8
	D_X5_8
	D_X6_8
	D_X7_8
	D_TLS_8    = 67
	D_NONE_8   = 68
	D_BRANCH_8 = 69
	D_EXTERN_8 = 70
	D_STATIC_8 = 71
	D_AUTO_8   = 72
	D_PARAM_8  = 73
	D_CONST_8  = 74
	D_FCONST_8 = 75
	D_SCONST_8 = 76
	D_ADDR_8   = 77 + iota - 50
	D_INDIR_8
	D_CONST2_8  = D_INDIR_8 + D_INDIR_8
	T_TYPE_8    = 1 << 0
	T_INDEX_8   = 1 << 1
	T_OFFSET_8  = 1 << 2
	T_FCONST_8  = 1 << 3
	T_SYM_8     = 1 << 4
	T_SCONST_8  = 1 << 5
	T_OFFSET2_8 = 1 << 6
	T_GOTYPE_8  = 1 << 7
	REGARG_8    = -1
	REGRET_8    = D_AX_8
	FREGRET_8   = D_F0_8
	REGSP_8     = D_SP_8
	REGTMP_8    = D_DI_8
)